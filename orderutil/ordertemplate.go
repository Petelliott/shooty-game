package orderutil

import (
    "github.com/petelliott/shooty-game/game"
	"encoding/json"
)

// this function is nuts
func Order(
	options *json.RawMessage,
	name string,
	opt interface{},
	conf interface{},
	handler func (unit *game.Unit, world *game.World),
) game.Order {

	err := json.Unmarshal([]byte(*options), opt)
	if err != nil {
		panic(err)
	}

	return func(unit *game.Unit) func(world *game.World) {
		err := json.Unmarshal([]byte(*unit.Utype.SupportedOrders[name]), conf)
		if err != nil {
			panic(err)
		}

		return func(world *game.World) {
			handler(unit, world)
		}
	}

}
