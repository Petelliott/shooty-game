package main

import (
	"github.com/petelliott/shooty-game/game"
	"github.com/gorilla/websocket"
	"github.com/petelliott/shooty-game/safews"
	"errors"
	"fmt"
	"net/http"
)


type Instance struct {
	Game *game.Game
	Teams map[game.Team]*safews.SafeConn
}

func NewInstance(g *game.Game) *Instance {
	i := Instance{g, make(map[game.Team]*safews.SafeConn)}
	return &i
}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func (i *Instance) NewTeam(team game.Team) (func(w http.ResponseWriter, r *http.Request), error) {
	if _, ok := i.Teams[team]; ok {
		return nil, errors.New(fmt.Sprintf("team %v is already in this instance", team))
	}

	wshandler := func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsupgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Failed to set websocket upgrade: %+v", err)
			return
		}

		sconn := safews.MakeSafe(conn)
		i.Teams[team] = sconn

		for {
			t, msg, err := sconn.ReadMessage()
			if err != nil {
				break
			}
			sconn.WriteMessage(t, msg)
		}
	}


	return wshandler, nil
}
