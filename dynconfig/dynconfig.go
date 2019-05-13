package dynconfig

import (
    "json"
    "plugin"
)

type Dynconfig struct {
    conf configFile
    plugindir string
    plugcache map[string]*Plugin
}

type unitClass struct {
    Graphic string `json:"graphic"`
    SupplyCap int `json:"supplycap"`
    SupportedOrders map[string]*json.RawMessage `json:"orders"`
}

type configFile struct {
    Units map[string]unitClass `json:"units"`
}
