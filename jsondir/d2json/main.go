package main

import (
	"github.com/petelliott/shooty-game/jsondir"
	"encoding/json"
	"os"
)

func main() {
	jd := jsondir.Jdir(os.Args[1], jsondir.DefaultEncoder())
	b, err := json.MarshalIndent(jd, "", "    ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
	os.Stdout.Write([]byte{'\n'})
}
