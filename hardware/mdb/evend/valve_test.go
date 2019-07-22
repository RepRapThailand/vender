package evend

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/temoto/vender/hardware/mdb"
	"github.com/temoto/vender/state"
)

func TestValve(t *testing.T) {
	t.Parallel()

	ctx, g := state.NewTestContext(t, `
engine { inventory {
	stock "water" { check=false hw_rate = 0.6 min = 500 }
}}`)
	mock := mdb.MockFromContext(ctx)
	defer mock.Close()
	go mock.Expect([]mdb.MockR{
		{"c0", ""},
		{"c1", "011810000a0000c8001fff01050a32640000000000000000000000"},

		{"c411", "17"},
		{"c51049", ""},

		{"c3", "44"},
		{"c3", "04"},
		{"c3", ""},
		{"c20136", ""},
		{"c3", "10"},
		{"c3", ""},
	})
	d := new(DeviceValve)
	// TODO make small delay default in tests
	d.dev.DelayIdle = 1
	d.dev.DelayNext = 1
	d.dev.DelayReset = 1
	err := d.Init(ctx)
	if err != nil {
		t.Fatalf("Init err=%v", err)
	}

	g.Engine.TestDo(t, ctx, "mdb.evend.valve_get_temp_hot")
	assert.Equal(t, uint8(23), uint8(d.tempHot.Get()))

	g.Engine.TestDo(t, ctx, "mdb.evend.valve_set_temp_hot(73)")

	water, err := g.Inventory.Get("water")
	require.NoError(t, err)
	initial := rand.Int31() - 10000 // check=false may go negative, not error
	t.Logf("water before=%d", initial)
	water.Set(initial)
	g.Engine.TestDo(t, ctx, "add.water_hot(90)")
	assert.Equal(t, initial-90, water.Value())
}
