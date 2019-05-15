package game

import (
    "encoding/json"
)

type Team string

const (
    RED Team = "red"
    BLUE Team = "blue"
)

type UnitClass struct {
    name string
    Graphic string `json:"graphic"`
    SupplyCap int `json:"supplycap"`
    SupportedOrders map[string]*json.RawMessage `json:"orders"`
}

func (u *UnitClass) Spawn(supplies int, x, y int, team Team) Unit {
    return Unit{u, x, y, team, supplies, nil}
}

func (u *UnitClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.name)
}

type Unit struct {
    utype *UnitClass `json:"utype"`
    x int `json:"x"`
    y int `json:"y"`
    team Team `json:"team"`
    Supplies int `json:"supplies"`
    currentOrder func()
}
