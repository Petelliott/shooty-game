package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "os"
    "github.com/petelliott/shooty-game/game"
    "github.com/petelliott/shooty-game/dynconfig"
    "github.com/petelliott/shooty-game/jsonconfig"
)

func main() {
    dc := dynconfig.Open(os.Args[1])
    jc := jsonconfig.Open(os.Args[2])

    game.NewStandardGame(&dc, &jc, 59)

    router := gin.Default()

    config := cors.DefaultConfig()
	config.AllowAllOrigins = true
    router.Use(cors.New(config))

    router.Run(":8049")
}
