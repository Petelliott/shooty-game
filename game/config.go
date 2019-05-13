package game

import (
    "encoding/json"
)

type Config interface {
    NewOrder(name string, options *json.RawMessage) Order
}
