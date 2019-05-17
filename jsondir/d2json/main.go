package main

import (
	"github.com/petelliott/shooty-game/jsondir"
	"encoding/json"
	"os"
)

func main() {
	b, _ := json.MarshalIndent(jsondir.Jdir(os.Args[1]), "", "    ")
	os.Stdout.Write(b)
	os.Stdout.Write([]byte{'\n'})
}
