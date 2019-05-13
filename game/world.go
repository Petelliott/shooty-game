package game


type World struct {
    Map Map `json:"map"`
    Units  []Unit `json:"units"`
}

func (world *World) Step() {
    for _, unit := range world.Units {
        unit.currentOrder()
    }
}

func (world *World) Order(unitno int, Order order) {
    Units[unitno] = func(unit *Unit) { order(Unists[unitno], world) }
}

