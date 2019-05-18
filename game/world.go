package game


type World struct {
    Map Map `json:"map"`
    Units  []Unit `json:"units"`
}

func (world *World) Step() {
	//TODO: random order of turns maybe?
    for _, unit := range world.Units {
        unit.currentOrder()
    }
}

func (world *World) Order(unitno int, order Order) {
	o := order(&world.Units[unitno])
    world.Units[unitno].currentOrder = func() { o(world) }
}

func (world *World) Spawn(u Unit) {
    world.Units = append(world.Units, u)
}

