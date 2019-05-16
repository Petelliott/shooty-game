package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/petelliott/shooty-game/dynconfig"
	"github.com/petelliott/shooty-game/game"
	"github.com/petelliott/shooty-game/jsonconfig"
	"os"
)

func main() {
	dc := dynconfig.Open(os.Args[1])
	jc := jsonconfig.Open(os.Args[2])

	g := game.NewGame(&dc, &jc, game.NewMap(10))
	inst := NewInstance(g)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	redhandler, _ := inst.NewTeam(game.RED)
	bluehandler, _ := inst.NewTeam(game.BLUE)

	router.GET("/ws/red", func(c *gin.Context) {
		redhandler(c.Writer, c.Request)
	})

	router.GET("/ws/blue", func(c *gin.Context) {
		bluehandler(c.Writer, c.Request)
	})

	router.Run(":8049")
}
