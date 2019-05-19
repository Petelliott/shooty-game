package jsondir

import (
	"io/ioutil"
	"fmt"
	"os"
	"path"
	"encoding/json"
)

type JsonDir struct {
	dir string
	encoder FileEncoder
}

func Jdir(dir string, encoder FileEncoder) JsonDir {
	return JsonDir{dir, encoder}
}


func (d JsonDir) MarshalJSON() ([]byte, error) {
	f, err := os.Stat(d.dir)
	if err != nil {
		return nil, fmt.Errorf("Error stating JsonDir: %v", err)
	}

	if f.Mode().IsRegular() {
		return d.encoder.Encode(d.dir)
	} else if f.Mode().IsDir() {
		files, err := ioutil.ReadDir(d.dir)
		if err != nil {
			return nil, fmt.Errorf("Error reading dir of JsonDir: %v", err)
		}

		m := make(map[string]*json.RawMessage)
		for _, sf := range files {
			data, err := json.Marshal(Jdir(path.Join(d.dir, sf.Name()), d.encoder))
			if err != nil {
				return nil, err
			}

			m[d.encoder.GetKey(sf.Name())] = (*json.RawMessage)(&data)
		}

		return json.Marshal(m)
	} else {
		return nil, fmt.Errorf("file %v is not a regular file or directory",
		err)
	}
}
