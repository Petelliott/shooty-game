package jsondir

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"path"
	"os"
	"fmt"
	"bytes"
)



func TestJsonDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "sgtest")
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}

	// cleanup
	defer os.RemoveAll(dir)

	err = os.MkdirAll(path.Join(dir, "a/a"), 0766)
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}

	err = ioutil.WriteFile(path.Join(dir, "a/a/a.json"), []byte(`{"x": 5, "y": 6}`), 0666)
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}

	err = ioutil.WriteFile(path.Join(dir, "b.json"), []byte(`7`), 0666)
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}

	err = ioutil.WriteFile(path.Join(dir, "plain"), []byte(`blah`), 0666)
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}


	jd := Jdir(dir, DefaultEncoder())

	b, err := json.Marshal(jd)
	if err != nil {
		fmt.Println(err)
		t.Error()
		return
	}

	if !bytes.Equal(b, []byte(`{"a":{"a":{"a":{"x":5,"y":6}}},"b":7,"plain":"blah"}`)) {
		fmt.Printf("got: '%s'", string(b))
		t.Error()
	}
}
