package game

type Order func(unit *Unit) (func(world *World))

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
