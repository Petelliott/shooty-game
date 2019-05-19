package main

import (
	"github.com/petelliott/shooty-game/jsondir"
	"encoding/json"
	"os"
	"runtime/debug"
)

func main() {
	b, err := json.MarshalIndent(jsondir.Jdir(os.Args[1]), "", "    ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
	os.Stdout.Write([]byte{'\n'})
}
