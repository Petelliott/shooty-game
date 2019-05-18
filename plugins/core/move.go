package main

import (
	"github.com/petelliott/shooty-game/game"
	"encoding/json"
	"math"
)

type MoveOptions struct {
	Destination game.Position `json:"destination"`
}

type MoveUnitConf struct {
	Speed		int `json:"speed"`
	Supplies	int `json:"supplies`
}

func Move(options *json.RawMessage) game.Order {
	var opt MoveOptions
	err := json.Unmarshal([]byte(*options), &opt)
	if err != nil {
		panic(err)
	}

	return func(unit *game.Unit) func(world *game.World) {
		var uc MoveUnitConf
		err := json.Unmarshal([]byte(*unit.Utype.SupportedOrders["core:Move"]), &uc)
		if err != nil {
			panic(err)
		}

		return func (world *game.World) {
			if unit.Supplies >= uc.Supplies {
				dx := float64(opt.Destination.X - unit.X)
				dy := float64(opt.Destination.Y - unit.Y)

				hyp := math.Sqrt(dx*dx + dy*dy)
				frac := float64(uc.Speed) / hyp

				unit.X += int(math.Round(frac*dx))
				unit.Y += int(math.Round(frac*dy))

				unit.Supplies -= uc.Supplies
			}
			// TODO some kind of alert to the user?
		}
	}
}
