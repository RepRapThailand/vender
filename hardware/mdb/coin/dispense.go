package coin

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/juju/errors"
	"github.com/temoto/vender/currency"
	"github.com/temoto/vender/engine"
	"github.com/temoto/vender/hardware/mdb"
	"github.com/temoto/vender/hardware/money"
	"github.com/temoto/vender/state"
)

// High-level coin dispense command. Handles:
// - built-in payout or "manual dispense" using expend strategy
// - give smallest amount >= requested
func (self *CoinAcceptor) NewDispenseSmart(requestAmount currency.Amount, over bool, success *currency.NominalGroup) engine.Doer {
	const tag = "mdb.coin.dispense-smart"

	return engine.Func{Name: tag, F: func(ctx context.Context) error {
		var err error
		var successAmount currency.Amount
		self.dev.Log.Debugf("%s requested=%s", tag, requestAmount.FormatCtx(ctx))
		leftAmount := requestAmount // save original requested amount for logs
		success.SetValid(self.nominals[:])
		if leftAmount == 0 {
			return nil
		}

		// === Try smart manual dispense
		if true { // FIXME DISPENSE is not completely implemented
			err = self.dispenseSmartManual(ctx, leftAmount, success)
			if err != nil {
				return errors.Annotate(err, tag)
			}
			successAmount = success.Total()
			if successAmount == requestAmount {
				return nil
			} else if successAmount > requestAmount {
				panic("code error")
			}
			leftAmount = requestAmount - successAmount
		}

		// === Fallback to PAYOUT
		self.dev.Log.Errorf("%s fallback to PAYOUT left=%s", tag, leftAmount.FormatCtx(ctx))
		err = self.NewPayout(leftAmount, success).Do(ctx)
		if err != nil {
			return errors.Annotate(err, tag)
		}
		successAmount = success.Total()
		if successAmount == requestAmount {
			return nil
		} else if successAmount > requestAmount {
			panic("code error")
		}
		leftAmount = requestAmount - successAmount

		// === Not enough coins for exact payout
		if !over {
			self.dev.Log.Errorf("%s not enough coins for exact payout and over-compensate disabled in request", tag)
			return currency.ErrNominalCount
		}
		// === Try to give a bit more
		// TODO telemetry
		successAmount = success.Total()
		self.dev.Log.Errorf("%s dispensed=%s < requested=%s debt=%s",
			tag, successAmount.FormatCtx(ctx), requestAmount.FormatCtx(ctx), leftAmount.FormatCtx(ctx))
		config := state.GetConfig(ctx)
		if leftAmount <= currency.Amount(config.Money.ChangeOverCompensate) {
			return self.NewDispenseLeastOver(leftAmount, success).Do(ctx)
		}
		return currency.ErrNominalCount
	}}
}

func (self *CoinAcceptor) NewDispenseLeastOver(requestAmount currency.Amount, success *currency.NominalGroup) engine.Doer {
	const tag = "mdb.coin.dispense-over"

	return engine.Func{Name: tag, F: func(ctx context.Context) error {
		var err error
		leftAmount := requestAmount

		nominals := self.SupportedNominals()
		sort.Slice(nominals, func(i, j int) bool { return nominals[i] < nominals[j] })

		for _, nominal := range nominals {
			namt := currency.Amount(nominal)

			// Round up `leftAmount` to nearest multiple of `nominal`
			payoutAmount := leftAmount + namt - 1 - (leftAmount-1)%namt

			self.dev.Log.Debugf("%s request=%s left=%s trying nominal=%s payout=%s",
				tag, requestAmount.FormatCtx(ctx), leftAmount.FormatCtx(ctx), namt.FormatCtx(ctx), payoutAmount.FormatCtx(ctx))
			payed := success.Copy()
			payed.Clear()
			// TODO use DISPENSE
			err = self.NewPayout(payoutAmount, payed).Do(ctx)
			success.AddFrom(payed)
			payedAmount := payed.Total()
			// self.dev.Log.Debugf("%s dispense err=%v", tag, err)
			if err != nil {
				return errors.Annotate(err, tag)
			}
			if leftAmount <= payedAmount {
				return nil
			}
			leftAmount -= payedAmount
		}
		return errors.Annotate(currency.ErrNominalCount, tag)
	}}
}

func (self *CoinAcceptor) dispenseSmartManual(ctx context.Context, amount currency.Amount, success *currency.NominalGroup) error {
	const tag = "mdb.coin.dispense-smart/manual"
	var err error

	self.DoTubeStatus.Do(ctx)
	tubeCoins := self.Tubes()
	if tubeCoins.Total() < amount {
		return nil
	}

	config := state.GetConfig(ctx)
	_ = config
	// TODO read preferred strategy from config
	strategy := currency.NewExpendLeastCount()
	if !strategy.Validate() {
		self.dev.Log.Errorf("%s config strategy=%v is not available, using fallback", tag, strategy)
		strategy = currency.NewExpendLeastCount()
		if !strategy.Validate() {
			panic("code error fallback coin strategy validate")
		}
	}

	ng := new(currency.NominalGroup)
	ng.SetValid(self.nominals[:])
	if err = tubeCoins.Withdraw(ng, amount, strategy); err != nil {
		// TODO telemetry
		self.dev.Log.Errorf("%s failed to calculate NominalGroup for dispense mode", tag)
		return nil
	}

	err = self.dispenseGroup(ctx, ng, success)
	self.dev.Log.Debugf("%s success=%s", tag, success.String())
	return errors.Annotate(err, tag)
}

func (self *CoinAcceptor) dispenseGroup(ctx context.Context, request, success *currency.NominalGroup) error {
	const tag = "mdb.coin.dispense-group"

	return request.Iter(func(nominal currency.Nominal, count uint) error {
		self.dev.Log.Debugf("%s n=%s c=%d", tag, currency.Amount(nominal).FormatCtx(ctx), count)
		if count == 0 {
			return nil
		}
		err := self.NewDispense(nominal, uint8(count)).Do(ctx)
		if err != nil {
			self.dev.Log.Errorf("%s nominal=%s count=%d err=%v", tag, currency.Amount(nominal).FormatCtx(ctx), count, err)
			return errors.Annotate(err, tag)
		}
		success.Add(nominal, count)
		return nil
	})
}

// MDB command DISPENSE (0d)
func (self *CoinAcceptor) NewDispense(nominal currency.Nominal, count uint8) engine.Doer {
	const tag = "mdb.coin.dispense"

	command := func(ctx context.Context) error {
		if count > 15 { // count must fit into 4 bits
			panic(fmt.Sprintf("code error %s count=%d > 15", tag, count))
		}
		coinType := self.nominalCoinType(nominal)
		if coinType < 0 {
			return errors.Errorf("%s not supported for nominal=%s", tag, currency.Amount(nominal).FormatCtx(ctx))
		}

		request := mdb.MustPacketFromBytes([]byte{0x0d, (count << 4) + uint8(coinType)}, true)
		err := self.dev.Tx(request).E
		return errors.Annotate(err, tag)
	}
	pollFun := func(p mdb.Packet) (bool, error) {
		bs := p.Bytes()
		if len(bs) == 0 {
			return true, nil
		}
		pi, _ := self.parsePollItem(bs[0], 0)
		// self.dev.Log.Debugf("%s poll=%x parsed=%v", tag, bs, pi)
		switch pi.Status {
		case money.StatusBusy:
			return false, nil
		case money.StatusError, money.StatusFatal: // tube jam, etc
			return true, pi.Error
		}
		return true, errors.Errorf("unexpected response=%x", bs)
	}

	return engine.Func{Name: tag, F: func(ctx context.Context) error {
		var err error
		// TODO  avoid double mutex acquire
		if err = self.DoTubeStatus.Do(ctx); err != nil {
			return errors.Annotate(err, tag)
		}
		tubesBefore := self.Tubes()
		var countBefore uint
		if countBefore, err = tubesBefore.Get(nominal); err != nil {
			return errors.Annotate(err, tag)
		} else if countBefore < uint(count) {
			err = currency.ErrNominalCount
			return errors.Annotate(err, tag)
		}

		self.pollmu.Lock()
		defer self.pollmu.Unlock()

		if err = command(ctx); err != nil {
			return errors.Annotate(err, tag)
		}
		totalTimeout := self.dispenseTimeout * time.Duration(count)
		if err = self.dev.NewPollLoop(tag, self.dev.PacketPoll, totalTimeout, pollFun).Do(ctx); err != nil {
			return errors.Annotate(err, tag)
		}

		if err = self.DoTubeStatus.Do(ctx); err != nil {
			return errors.Annotate(err, tag)
		}
		self.CommandExpansionSendDiagStatus(nil)
		tubesAfter := self.Tubes()
		var countAfter uint
		if countAfter, err = tubesAfter.Get(nominal); err != nil {
			return errors.Annotate(err, tag)
		}

		diff := int(countBefore) - int(countAfter)
		if diff != int(count) {
			return errors.Errorf("%s nominal=%s requested=%d diff=%d", tag, currency.Amount(nominal).FormatCtx(ctx), count, diff)
		}
		return nil
	}}
}

// MDB command PAYOUT (0f02)
func (self *CoinAcceptor) NewPayout(amount currency.Amount, success *currency.NominalGroup) engine.Doer {
	const tag = "mdb.coin.payout"
	self.dev.Log.Debugf("%s sf=%v amount=%s", tag, self.scalingFactor, amount.Format100I())
	arg := amount / currency.Amount(self.scalingFactor)

	doPayout := engine.Func{Name: tag + "/command", F: func(ctx context.Context) error {
		request := mdb.MustPacketFromBytes([]byte{0x0f, 0x02, byte(arg)}, true)
		err := self.dev.Tx(request).E
		return errors.Annotate(err, tag)
	}}
	doStatus := engine.Func{Name: tag + "/status", F: func(ctx context.Context) error {
		r := self.dev.Tx(packetPayoutStatus)
		if r.E != nil {
			return errors.Annotate(r.E, tag)
		}
		for i, count := range r.P.Bytes() {
			nominal := self.nominals[i]
			if err := success.Add(nominal, uint(count)); err != nil {
				return errors.Annotate(err, tag)
			}
		}
		return nil
	}}

	return engine.NewSeq(tag).
		Append(doPayout).
		Append(self.dev.NewPollLoop(tag, packetPayoutPoll, self.dispenseTimeout*4, payoutPollFun)).
		Append(doStatus)
}

// 0FH 04H PAYOUT VALUE POLL
// - [Response is] 1 byte scaled amount of paid out change since [... payout/poll]
// - An 00H response indicates no coins were paid out since [... payout/poll]
// - An ACK only indicates that the change payout is finished.
//   This should be followed by the PAYOUT STATUS command (0FH-03H) to obtain the complete payout data.
func payoutPollFun(p mdb.Packet) (bool, error) {
	if p.Len() == 0 {
		return true, nil
	}
	return false, nil
}
