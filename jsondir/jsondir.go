package jsondir

import (
	"io/ioutil"
	"fmt"
	"os"
	"path"
	"encoding/json"
)

type JsonDir string

func Jdir(dir string) JsonDir {
	return JsonDir(dir)
}


func (d JsonDir) MarshalJSON() ([]byte, error) {
	dir := string(d)

	f, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("Error stating JsonDir: %v", err)
	}

	encoder := DefaultEncoder()
	if f.Mode().IsRegular() {
		return encoder.Encode(dir)
	} else if f.Mode().IsDir() {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("Error reading dir of JsonDir: %v", err)
		}

		m := make(map[string]*json.RawMessage)
		for _, sf := range files {
			data, err := json.Marshal(Jdir(path.Join(dir, sf.Name())))
			if err != nil {
				return nil, err
			}

			m[encoder.GetKey(sf.Name())] = (*json.RawMessage)(&data)
		}

		return json.Marshal(m)
	} else {
		return nil, fmt.Errorf("file %v is not a regular file or directory",
		err)
	}
}
