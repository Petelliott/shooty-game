package game

import (
    "encoding/json"
)

type UnitClass struct {
    name string
    Graphic string `json:"graphic"`
    SupplyCap int `json:"supplycap"`
    SupportedOrders map[string]*json.RawMessage `json:"orders"`
}

func (u *UnitClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.name)
}

type Unit struct {
    utype *UnitClass `json:"utype"`
    x int `json:"x"`
    y int `json:"y"`
    Supplies int `json:"supplies"`
    currentOrder func()
}
