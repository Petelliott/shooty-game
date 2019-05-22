package game

import (
	"testing"
	"encoding/json"
)

type testOrderConfig struct {}
type testUnitConfig struct {}

func (oc *testOrderConfig) NewOrder(name string, options *json.RawMessage) Order {
	if name == "a" {
		return func(unit *Unit) (func(world *World)) {
			return func(world *World) {
				unit.X = 1000
				unit.Y = 10
			}
		}
	} else if name == "b" {
		return func(unit *Unit) (func(world *World)) {
			return func(world *World) {
				unit.X = 6
				unit.Y = 19
				unit.Supplies = 65
			}
		}
	} else {
		return nil
	}
}

var uca = UnitClass{"a", "", 1000, nil}
var ucb = UnitClass{"b", "", 2000, nil}
var ucbase = UnitClass{"core:Base", "", 2500, nil}

func (uc *testUnitConfig) GetUnitClass(name string) *UnitClass {
	if name == "a" {
		return &uca
	} else if name == "b" {
		return &ucb
	} else if name == "core:Base" {
		return &ucbase
	} else {
		return nil
	}
}

func TestGame1(t *testing.T) {
	oc := testOrderConfig{}
	uc := testUnitConfig{}
	g := NewGame(&oc, &uc, NewMap(10))

	g.AddTeam(5,5, RED)
	g.AddTeam(10,10, BLUE)

	options := json.RawMessage([]byte{})
	g.World.Order(0, g.OrderConfig.NewOrder("a", &options))
	g.World.Order(1, g.OrderConfig.NewOrder("b", &options))

	g.World.Step()

	rt := g.World.Units[0]
	bt := g.World.Units[1]

	if rt.Utype != &ucbase { t.Error() }
	if rt.X != 1000 { t.Error() }
	if rt.Y != 10 { t.Error() }
	if rt.Team != RED { t.Error() }

	if bt.Utype != &ucbase { t.Error() }
	if bt.X != 6 { t.Error() }
	if bt.Y != 19 { t.Error() }
	if bt.Supplies != 65 { t.Error() }
	if bt.Team != BLUE { t.Error() }

}
