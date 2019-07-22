package inventory

import (
	"context"
	"sync"

	"github.com/temoto/errors"
	"github.com/temoto/vender/engine"
	engine_config "github.com/temoto/vender/engine/config"
	"github.com/temoto/vender/helpers"
	"github.com/temoto/vender/log2"
	"github.com/temoto/vender/state/persist"
)

var (
	ErrStockLow = errors.New("Stock is too low")
)

type Inventory struct {
	persist.Persist
	engine *engine.Engine
	log    *log2.Log
	mu     sync.RWMutex
	ss     map[string]*Stock
}

func (self *Inventory) Init(ctx context.Context, c *engine_config.Inventory, engine *engine.Engine) error {
	self.log = log2.ContextValueLogger(ctx)
	self.engine = engine

	self.mu.Lock()
	defer self.mu.Unlock()
	errs := make([]error, 0)
	self.ss = make(map[string]*Stock, 32)
	for _, stockConfig := range c.Stocks {
		if _, ok := self.ss[stockConfig.Name]; ok {
			errs = append(errs, errors.Errorf("stock=%s already registered", stockConfig.Name))
			continue
		}

		stock, err := NewStock(stockConfig, self.engine)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		self.ss[stock.Name] = stock
	}

	return helpers.FoldErrors(errs)
}

func (self *Inventory) EnableAll()  { self.Iter(func(s *Stock) { s.Enable() }) }
func (self *Inventory) DisableAll() { self.Iter(func(s *Stock) { s.Disable() }) }

func (self *Inventory) Get(name string) (*Stock, error) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	if s, ok := self.ss[name]; ok {
		return s, nil
	}
	return nil, errors.Errorf("stock=%s is not registered", name)
}

func (self *Inventory) Iter(fun func(s *Stock)) {
	self.mu.Lock()
	for _, stock := range self.ss {
		fun(stock)
	}
	self.mu.Unlock()
}
