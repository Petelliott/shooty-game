package jsonconfig

import (
    "github.com/petelliott/shooty-game/game"
    "io/ioutil"
    "encoding/json"
)

type Jsonconfig struct {
    conf configfile
}

type configfile struct {
    units map[string]*game.UnitClass `json:"units"`
}

func Open(path string) Jsonconfig {
    plan, _ := ioutil.ReadFile(path)
    var data Jsonconfig
    err := json.Unmarshal(plan, &data)
    if err != nil {
        panic(err)
    }
    return data
}

func (jc *Jsonconfig) GetUnitClass(name string) *game.UnitClass {
    uc := jc.conf.units[name]
    uc.Name = name
    return uc
}
