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
	DoneTurn map[game.Team]bool
}


func NewInstance(g *game.Game) *Instance {
	i := Instance{g, make(map[game.Team]*safews.SafeConn), make(map[game.Team]bool)}
	return &i
}


func (i *Instance) EndTurn(team game.Team) {
	i.DoneTurn[team] = true

	for _, d := range i.DoneTurn {
		if d == false {
			return
		}
	}

	// everyone is done this turn now

	// run the commands for this turn
	i.Game.World.Step()

	for t, c := range i.Teams {
		i.DoneTurn[t] = false
		// send everyone the state
		c.WriteJSON(i.Game.World.Units)
	}
}


var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}


func (i *Instance) NewTeam(team game.Team) (func(w http.ResponseWriter, r *http.Request), error) {
	if _, ok := i.Teams[team]; ok {
		return nil, errors.New(fmt.Sprintf("team %v is already in this instance", team))
	}

	i.DoneTurn[team] = false

	wshandler := func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsupgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("Failed to set websocket upgrade: %+v", err)
			return
		}

		sconn := safews.MakeSafe(conn)
		i.Teams[team] = sconn

		// send the World
		sconn.WriteJSON(i.Game.World)

		for {
			var msg Msg
			err := sconn.ReadJSON(&msg)
			if err != nil {
				break
			}
			i.Game.OrderConfig.NewOrder(msg.Order, msg.Options)
			i.EndTurn(team)
		}
	}


	return wshandler, nil
}
