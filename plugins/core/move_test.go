package main

import (
	"testing"
	"encoding/json"
	"github.com/petelliott/shooty-game/game"
	"fmt"
)

func TestMove(t *testing.T) {
	args := json.RawMessage(`{"destination": {"x": 10, "y": 0}}`)
	options := json.RawMessage(`{"speed": 1, "supplies": 1}`)
	uca := game.UnitClass{"a", "", 1000, map[string]*json.RawMessage{"core:Move": &options}}

	unit := uca.Spawn(3, 0, 0, game.RED)

	f := Move(&args)

	f(&unit)(nil)
	if unit.X != 1 || unit.Y != 0 || unit.Supplies != 2 {
		fmt.Println(unit)
		t.Error()
	}

	f(&unit)(nil)
	if unit.X != 2 || unit.Y != 0 || unit.Supplies != 1 {
		fmt.Println(unit)
		t.Error()
	}

	f(&unit)(nil)
	if unit.X != 3 || unit.Y != 0 || unit.Supplies != 0 {
		fmt.Println(unit)
		t.Error()
	}

	f(&unit)(nil)
	if unit.X != 3 || unit.Y != 0 || unit.Supplies != 0 {
		fmt.Println(unit)
		t.Error()
	}

	f(&unit)(nil)
	if unit.X != 3 || unit.Y != 0 || unit.Supplies != 0 {
		fmt.Println(unit)
		t.Error()
	}
}
