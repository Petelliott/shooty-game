package game

import "encoding/json"

type Map struct {
	Tiles          []byte          `json:"tiles"`
	size           int             `json:"size"`
	TileDescriptor map[byte]string `json:"tileDescription"`
}

func NewMap(size int) Map {
	tiles := make([]byte, size*size)
	td := map[byte]string{
		0: "ground",
	}
	return Map{tiles, size, td}
}

func (m *Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m)
}
