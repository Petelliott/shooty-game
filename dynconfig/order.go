package dynconfig

import (
    "encoding/json"
    "github.com/petelliott/shooty-game/game"
    "strings"
    "plugin"
)

type OrderDesignator func(options *json.RawMessage) game.Order

func (conf *Dynconfig) NewOrder(name string, options *json.RawMessage) game.Order {
    spec := strings.Split(name, ":")
    pkg := spec[0]
    command := spec[1]

    plug, err := plugin.Open(conf.plugindir + pkg + ".so")
    if err != nil {
        panic(err)
    }

    order, err := plug.Lookup(command)
    if err != nil {
        panic(err)
    }

    return order.(OrderDesignator)(options)
}
