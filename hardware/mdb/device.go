package mdb

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/juju/errors"
	"github.com/temoto/vender/engine"
	"github.com/temoto/vender/helpers/atomic_clock"
	"github.com/temoto/vender/log2"
)

const ErrCodeNone int32 = -1

const (
	DefaultDelayAfterReset  = 500 * time.Millisecond
	DefaultDelayBeforeReset = 0
	DefaultDelayIdle        = 700 * time.Millisecond
	DefaultDelayNext        = 200 * time.Millisecond
	DefaultDelayOffline     = 10 * time.Second
	DefaultIdleThreshold    = 30 * time.Second
)

type PacketError struct {
	E error
	P Packet
}

type Device struct { //nolint:maligned
	state   uint32 // atomic
	errCode int32  // atomic

	bus   *Bus
	cmdLk sync.Mutex // TODO explore if chan approach is better

	LastOk      *atomic_clock.Clock // last successful tx(), 0 at init, monotonic
	LastOff     *atomic_clock.Clock // last change from online to offline (MDB timeout), 0=online
	lastReset   *atomic_clock.Clock // last RESET attempt, 0 only at init, monotonic
	Log         *log2.Log
	Address     uint8
	Name        string
	ByteOrder   binary.ByteOrder
	PacketReset Packet
	PacketSetup Packet
	PacketPoll  Packet
	DoReset     engine.Doer
	DoInit      engine.Doer // likely Seq starting with DoReset

	DelayAfterReset  time.Duration
	DelayBeforeReset time.Duration
	DelayIdle        time.Duration
	DelayNext        time.Duration
	DelayOffline     time.Duration
	IdleThreshold    time.Duration

	SetupResponse Packet
}

func (self *Device) Init(bus *Bus, addr uint8, name string, byteOrder binary.ByteOrder) {
	self.cmdLk.Lock()
	defer self.cmdLk.Unlock()

	self.Address = addr
	self.ByteOrder = byteOrder
	self.Log = bus.Log
	self.bus = bus
	self.Name = name
	self.errCode = ErrCodeNone
	self.LastOk = atomic_clock.New(0)
	self.LastOff = atomic_clock.Now()
	self.lastReset = atomic_clock.New(0)

	if self.DelayAfterReset == 0 {
		self.DelayAfterReset = DefaultDelayAfterReset
	}
	if self.DelayBeforeReset == 0 {
		self.DelayBeforeReset = DefaultDelayBeforeReset
	}
	if self.DelayIdle == 0 {
		self.DelayIdle = DefaultDelayIdle
	}
	if self.DelayNext == 0 {
		self.DelayNext = DefaultDelayNext
	}
	if self.DelayOffline == 0 {
		self.DelayOffline = DefaultDelayOffline
	}
	if self.IdleThreshold == 0 {
		self.IdleThreshold = DefaultIdleThreshold
	}
	self.SetupResponse = Packet{}
	self.PacketReset = MustPacketFromBytes([]byte{self.Address + 0}, true)
	self.PacketSetup = MustPacketFromBytes([]byte{self.Address + 1}, true)
	self.PacketPoll = MustPacketFromBytes([]byte{self.Address + 3}, true)
	self.DoReset = engine.Func0{Name: fmt.Sprintf("mdb.%s.reset", self.Name), F: self.Reset}
	self.SetState(DeviceInited)
}

func (self *Device) ValidateErrorCode() error {
	value := atomic.LoadInt32(&self.errCode)
	if value == ErrCodeNone {
		return nil
	}
	return errors.Errorf("mdb.%s unhandled errorcode=%d", self.Name, value)
}

func (self *Device) ValidateOnline() error {
	st := self.State()
	if st.Online() {
		return nil
	}
	return errors.Errorf("mdb.%s state=%s offline duration=%v", self.Name, st.String(), atomic_clock.Since(self.LastOff))
}

func (self *Device) Probe(request Packet) error {
	self.cmdLk.Lock()
	defer self.cmdLk.Unlock()

	valid := false
	switch st := self.State(); st {
	case DeviceInited: // success path at reboot

	case DeviceInvalid:
		panic("code error Probe before Init")

	case DeviceOffline, DeviceError, DeviceReady:
		now := atomic_clock.Now()
		if age := now.Sub(self.LastOk); !valid && age >= 0 && age < self.DelayOffline {
			valid = true
		} else if age := now.Sub(self.LastOff); !valid && age >= 0 && age < self.DelayOffline {
			valid = true
		}
		if valid {
			switch st {
			case DeviceOffline:
				return ErrOffline
			case DeviceReady:
				return nil
			case DeviceError:
				if err := self.ValidateErrorCode(); err != nil {
					return err
				}
				return self.ValidateErrorCode()
			default:
				panic("code error Probe valid=true switch st")
			}
		} else {
			time.Sleep(self.DelayAfterReset)
		}
	}

	r := self.txInternal(self.PacketPoll)
	if r.E == nil {
		self.SetState(DeviceOnline)
		return nil
	}
	if errors.Cause(r.E) == ErrTimeout {
		self.SetState(DeviceOffline)
		return errors.Wrap(r.E, ErrOffline)
	}
	// TODO parse error code
	self.SetState(DeviceError)
	return errors.Annotate(r.E, "probe")
}

func (self *Device) Tx(request Packet) (r PacketError) {
	if err := self.ValidateErrorCode(); err != nil {
		// self.Log.Errorf("TODO-ERRCODE %v", errors.ErrorStack(errors.Trace(err)))
		// TODO self.Reset()
	}
	return self.tx(request)
}

func (self *Device) DoSetup(ctx context.Context) error {
	self.cmdLk.Lock()
	defer self.cmdLk.Unlock()
	self.SetupResponse = *PacketEmpty
	r := self.tx(self.PacketSetup)
	if r.E != nil {
		return errors.Annotatef(r.E, "mdb.%s SETUP", self.Name)
	}
	self.SetupResponse = r.P
	return nil
}

func (self *Device) ErrorCode() int32 { return atomic.LoadInt32(&self.errCode) }
func (self *Device) SetErrorCode(c int32) {
	prev := atomic.SwapInt32(&self.errCode, c)
	if prev != ErrCodeNone {
		self.Log.Errorf("mdb.%s CRITICAL SetErrorCode overwrite previous=%d", self.Name, prev)
		// TODO tele
	}
	if c != ErrCodeNone {
		self.SetState(DeviceError)
		err := fmt.Errorf("mdb.%s errcode=%d", self.Name, c)
		self.bus.Error(err)
	}
}

func (self *Device) State() DeviceState       { return DeviceState(atomic.LoadUint32(&self.state)) }
func (self *Device) Ready() bool              { return self.State() == DeviceReady }
func (self *Device) SetState(new DeviceState) { atomic.StoreUint32(&self.state, uint32(new)) }
func (self *Device) SetReady()                { self.SetState(DeviceReady) }
func (self *Device) SetOnline()               { self.SetState(DeviceOnline) }

func (self *Device) Reset() error {
	self.cmdLk.Lock()
	defer self.cmdLk.Unlock()
	return self.locked_reset()
}

// Main purpose is to keep certain manufacturer devices "hot" to reduce useless POLL time.
// Incidentally this is also adequate place to put regular RESET attempts for offline devices.
func (self *Device) Keepalive(interval time.Duration, stopch <-chan struct{}) {
	wait := interval

	for {
		// TODO try and benchmark time.After vs NewTimer vs NewTicker
		// self.Log.Debugf("keepalive wait=%v", wait)
		if wait <= 0 {
			wait = 1
		}
		select {
		case <-stopch:
			return
		case <-time.After(wait):
		}
		if online := self.State().Online(); !online {
			self.cmdLk.Lock()
			wait, online, _ = self.locked_delayedReset()
			self.cmdLk.Unlock()
			if !online {
				continue
			}
		}

		self.cmdLk.Lock()
		// Could become offline during Lock()
		if self.State().Online() {
			okAge := atomic_clock.Since(self.LastOk)
			wait = interval - okAge
			// self.Log.Debugf("keepalive locked okage=%v wait=%v", okAge, wait)
			if wait <= 0 {
				// TODO validate err code?
				self.txInternal(self.PacketPoll)
				wait = interval
			}
		}
		self.cmdLk.Unlock()
	}
}

type PollFunc func(Packet) (stop bool, err error)

func (self *Device) NewPollLoop(tag string, request Packet, timeout time.Duration, fun PollFunc) engine.Doer {
	tag += "/poll-loop"
	return engine.Func{Name: tag, F: func(ctx context.Context) error {
		var r PacketError
		tbegin := time.Now()

		self.cmdLk.Lock()
		defer self.cmdLk.Unlock()
		for {
			// TODO validate err code?
			r = self.tx(request)
			if r.E != nil {
				return errors.Annotate(r.E, tag)
			}
			stop, err := fun(r.P)
			if err == nil && stop { // success
				return nil
			} else if err == nil && !stop { // try again
				if timeout == 0 {
					return errors.Errorf("tag=%s timeout=0 invalid", tag)
				}
				time.Sleep(self.DelayNext)
				if time.Since(tbegin) > timeout {
					return errors.Timeoutf(tag)
				}
				continue
			}

			return errors.Annotate(err, tag)
		}
	}}
}

// Used by tests to avoid waiting.
func (self *Device) XXX_FIXME_SetAllDelays(d time.Duration) {
	self.DelayIdle = d
	self.DelayNext = d
	self.DelayBeforeReset = d
	self.DelayAfterReset = d
	self.DelayOffline = d
}

// Returns retry delay and online status.
func (self *Device) locked_delayedReset() (time.Duration, bool, error) {
	// LastOff could be updated during Lock()
	if self.State().Online() {
		return 0, true, nil
	}
	resetAge := atomic_clock.Since(self.lastReset)
	if resetAge < self.DelayOffline { // don't RESET too often
		return self.DelayOffline - resetAge, false, nil
	}

	if err := self.locked_reset(); err != nil {
		return self.DelayOffline, false, err
	}
	return 0, true, nil
}

// cmdLk used to ensure no concurrent commands during delay
func (self *Device) locked_reset() error {
	tag := fmt.Sprintf("mdb.%s", self.Name)
	self.LastOff.SetNowIfZero() // consider device offline from now till successful response
	self.SetState(DeviceInited)
	time.Sleep(self.DelayBeforeReset)
	r := self.txInternal(self.PacketReset)
	self.lastReset.SetNow()
	atomic.StoreInt32(&self.errCode, ErrCodeNone)
	if r.E != nil {
		err := r.E
		if errors.Cause(err) == ErrTimeout {
			self.SetState(DeviceOffline)
			// err = errors.Annotatef(err, "%s is offline", tag)
			err = errors.Wrap(err, ErrOffline)
		} else {
			// TODO remove log here when ensured that error is logged in all callers
			// - Keepalive() ignores err
			// - Reset() returns err to caller without logging
			self.Log.Errorf("%s RESET err=%s", tag, errors.ErrorStack(err))
		}
		err = errors.Annotatef(err, "%s RESET", tag)
		return err
	}
	self.Log.Infof("%s addr=%02x is working", tag, self.Address)
	time.Sleep(self.DelayAfterReset)
	return nil
}

func (self *Device) tx(request Packet) PacketError {
	if err := self.ValidateOnline(); err != nil {
		return PacketError{E: err}
	}
	return self.txInternal(request)
}

func (self *Device) txInternal(request Packet) (r PacketError) {
	st := self.State()
	if st == DeviceInvalid {
		r.E = errors.Errorf("mdb.%s state=invalid", self.Name)
		return
	}

	r.E = self.bus.Tx(request, &r.P)
	if r.E == nil {
		// self.Log.Debugf("mdb.%s since last ok %v", self.Name, atomic_clock.Since(self.LastOk))
		self.LastOk.SetNow()
		self.LastOff.Set(0)
		// Upgrade any state except Ready to Online
		// Ready->Online would loose calibration.
		if st != DeviceReady {
			self.SetOnline()
		}
		atomic.StoreInt32(&self.errCode, ErrCodeNone)
	} else {
		if st.Online() {
			self.SetErrorCode(0)
		}
		r.E = errors.Annotatef(r.E, "request=%x", request.Bytes())
		self.Log.Errorf("mdb.%s err=%v", self.Name, r.E)
	}
	self.Log.Debugf("mdb.%s tx %x -> ok=%t/timeout=%t state %s -> %s",
		self.Name, request.Bytes(), r.E == nil, errors.IsTimeout(errors.Cause(r.E)), st.String(), self.State().String())
	return
}

// "Idle mode" polling, runs forever until receive on `stopch`.
// Switches between fast/idle delays.
// Used by bill/coin devices.
type PollDelay struct {
	lastActive time.Time
	lastDelay  time.Duration
}

func (self *PollDelay) Delay(dev *Device, active bool, err bool, stopch <-chan struct{}) bool {
	delay := dev.DelayNext

	if err {
		delay = dev.DelayIdle
	} else if active {
		self.lastActive = time.Now()
	} else if self.lastDelay != dev.DelayIdle { // save time syscall while idle continues
		if time.Since(self.lastActive) > dev.IdleThreshold {
			delay = dev.DelayIdle
		}
	}
	self.lastDelay = delay

	select {
	case <-stopch:
		return false
	case <-time.After(delay):
		return true
	}
}
