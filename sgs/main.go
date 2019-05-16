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

	game.NewGame(&dc, &jc, game.NewMap(10))

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Run(":8049")
}
