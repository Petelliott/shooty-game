package orderutil

import (
	"testing"
	"github.com/petelliott/shooty-game/game"
	"encoding/json"
)

type testMoveOptions struct {
	Destination game.Position `json:"destination"`
}

type testMoveUnitConf struct {
	Speed    int `json:"speed"`
	Supplies int `json:"supplies`
}

func TestOrder(t *testing.T) {
	args := json.RawMessage(`{"destination": {"x": 10, "y": 10}}`)

	options := json.RawMessage(`{"speed": 1, "supplies": 1}`)
	uca := game.UnitClass{"a", "", 1000, map[string]*json.RawMessage{"core:Move": &options}}

	unit := uca.Spawn(1000, 0, 0, game.RED)

	var opt testMoveOptions
	var conf testMoveUnitConf
	f := Order(&args, "core:Move", &opt, &conf,
		func(unit *game.Unit, world *game.World) {
			if opt.Destination.X != 10 { t.Error() }
			if opt.Destination.Y != 10 { t.Error() }
			if conf.Speed != 1 { t.Error() }
			if conf.Supplies != 1 { t.Error() }
			unit.X = 500
			unit.Y = 600
		})

	f(&unit)(nil)

	if unit.X != 500 { t.Error() }
	if unit.Y != 600 { t.Error() }
}
