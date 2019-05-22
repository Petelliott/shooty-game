package game

type Game struct {
    OrderConfig OrderConfig
    UnitsConfig UnitsConfig
    World World
}

func NewGame(oc OrderConfig, uc UnitsConfig, m Map) *Game {
    game := Game{oc, uc, World{m, []Unit{}}}

    return &game
}

func (g *Game) AddTeam(x, y int, team Team) {
	g.World.Spawn(
		g.UnitsConfig.GetUnitClass("core:Base").Spawn(1000, x, y, team))
}
