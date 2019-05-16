package main

import (
	"encoding/json"
)

type Msg struct {
	IsReady bool				`json:"isready"`
	Order string				`json:"order"`
	Options *json.RawMessage	`json:"options"`
}

