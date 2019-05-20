package dynconfig

import (
	"testing"
	"encoding/json"
	"github.com/petelliott/shooty-game/game"
)


func TestDynconfig(t *testing.T) {
	dc := Open("../gamedir/plugins")
	target := json.RawMessage(`{"destination": {"x": 10, "y": 10}}`)
	order := dc.NewOrder("core:Move", &target)

	options := json.RawMessage(`{"speed": 1, "supplies": 1}`)
	uca := game.UnitClass{"a", "", 1000, map[string]*json.RawMessage{"core:Move": &options}}

	unit := uca.Spawn(1000, 0, 0, game.RED)

	order(&unit)(nil)

	if unit.X == 0 { t.Error() }
	if unit.Y == 0 { t.Error() }
	if unit.Supplies != 999 { t.Error() }
}
