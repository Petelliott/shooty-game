package game

import (
    "encoding/json"
)

type OrderConfig interface {
    NewOrder(name string, options *json.RawMessage) Order
}
