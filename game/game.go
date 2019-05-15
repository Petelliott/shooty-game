package game

type Game struct {
    OrderConfig OrderConfig
    UnitsConfig UnitsConfig
    World World
}

func NewStandardGame(oc OrderConfig, uc UnitsConfig, m Map) *Game {
    game := Game{oc, uc, World{m, []Unit{}}}

    game.World.Spawn(uc.GetUnitClass("base").Spawn(1000, 0, 0, RED))
    game.World.Spawn(uc.GetUnitClass("base").Spawn(1000, 5, 5, BLUE))

    return &game
}
