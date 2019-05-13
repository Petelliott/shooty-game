package game


type Unitid string

type Unit struct {
    utype Unitid `json:"utype"`
    x int `json:"x"`
    y int `json:"y"`
    Supplies int `json:"supplies"`
    currentOrder func()
}
