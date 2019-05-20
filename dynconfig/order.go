package dynconfig

import (
    "encoding/json"
    "github.com/petelliott/shooty-game/game"
    "strings"
    "plugin"
	"path"
)

func (conf *Dynconfig) NewOrder(name string, options *json.RawMessage) game.Order {
    spec := strings.Split(name, ":")
    pkg := spec[0]
    command := spec[1]

    plug, err := plugin.Open(path.Join(conf.plugindir, pkg + ".so"))
    if err != nil {
        panic(err)
    }

    order, err := plug.Lookup(command)
    if err != nil {
        panic(err)
    }

    return order.(func(options *json.RawMessage) game.Order)(options)
}
