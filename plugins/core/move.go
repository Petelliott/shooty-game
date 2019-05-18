package main

import (
	"encoding/json"
	"github.com/petelliott/shooty-game/game"
	"github.com/petelliott/shooty-game/orderutil"
	"math"
)

type MoveOptions struct {
	Destination game.Position `json:"destination"`
}

type MoveUnitConf struct {
	Speed    int `json:"speed"`
	Supplies int `json:"supplies`
}

func Move(options *json.RawMessage) game.Order {
	var opt MoveOptions
	var conf MoveUnitConf

	return orderutil.Order(options, "core:Move", &opt, &conf,
		func(unit *game.Unit, world *game.World) {
			if unit.Supplies >= conf.Supplies {
				dx := float64(opt.Destination.X - unit.X)
				dy := float64(opt.Destination.Y - unit.Y)

				hyp := math.Sqrt(dx*dx + dy*dy)
				frac := float64(conf.Speed) / hyp

				unit.X += int(math.Round(frac * dx))
				unit.Y += int(math.Round(frac * dy))

				unit.Supplies -= conf.Supplies
			}
			// TODO some kind of alert to the user?
		})
}
