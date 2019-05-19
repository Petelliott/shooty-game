package jsondir

import (
	"testing"
	"bytes"
	"io/ioutil"
	"os"
)

func TestAddKeyHandler(t *testing.T) {
	e := NewEncoder()

	e.AddKeyHandler(".a", func(path string) string {
		return path+"/a"
	})

	e.AddKeyHandler(".b", func(path string) string {
		return path+"/b"
	})

	// no file extention
	e.AddKeyHandler("", func(path string) string {
		return path+"/none"
	})

	if e.GetKey("test.a") != "test.a/a" {
		t.Error()
	}

	if e.GetKey("test.b") != "test.b/b" {
		t.Error()
	}

	if e.GetKey("test") != "test/none" {
		t.Error()
	}

	// test overwriting
	e.AddKeyHandler(".a", func(path string) string {
		return path+"/aa"
	})

	if e.GetKey("test.a") != "test.a/aa" {
		t.Error()
	}
}

func TestDefaultKeyHandler(t *testing.T) {
	e := NewEncoder()

	e.AddKeyHandler(".a", func(path string) string {
		return path+"/a"
	})

	if e.GetKey("test.asdf") != "test.asdf" {
		t.Error()
	}

	e.SetDefaultKeyHandler(func(path string) string {
		return path+"/def"
	})

	if e.GetKey("test.v") != "test.v/def" {
		t.Error()
	}

	// test that setting default does not reset handlers
	if e.GetKey("test.a") != "test.a/a" {
		t.Error()
	}
}

func TestAddValueHandler(t *testing.T) {
	e := NewEncoder()

	e.AddValueHandler(".a", func(path string) ([]byte, error) {
		return []byte("a"), nil
	})

	e.AddValueHandler(".b", func(path string) ([]byte, error) {
		return []byte("b"), nil
	})

	if d, err := e.Encode("test.a"); !bytes.Equal(d, []byte("a")) || err != nil {
		t.Error()
	}

	if d, err := e.Encode("test.b"); !bytes.Equal(d, []byte("b")) || err != nil {
		t.Error()
	}

	// test replaceing an extension
	e.AddValueHandler(".a", func(path string) ([]byte, error) {
		return []byte("aa"), nil
	})

	if d, err := e.Encode("test.a"); !bytes.Equal(d, []byte("aa")) || err != nil {
		t.Error()
	}
}

func TestSetDefaultValueHandler(t *testing.T) {
	e := NewEncoder()

	e.SetDefaultValueHandler(func(path string) ([]byte, error) {
		return []byte("def1"), nil
	})

	e.AddValueHandler(".a", func(path string) ([]byte, error) {
		return []byte("a"), nil
	})

	if d, err := e.Encode("test.cdx"); !bytes.Equal(d, []byte("def1")) || err != nil {
		t.Error()
	}

	if d, err := e.Encode("test"); !bytes.Equal(d, []byte("def1")) || err != nil {
		t.Error()
	}

	e.SetDefaultValueHandler(func(path string) ([]byte, error) {
		return []byte("def2"), nil
	})

	if d, err := e.Encode("test"); !bytes.Equal(d, []byte("def2")) || err != nil {
		t.Error()
	}
}

func TestDefaultEncoder(t *testing.T) {
	DefaultEncoder()

	// this test doesn't really need to do anything, all this behavior is tested otherwise
}

func withTestFile(t *testing.T, pat string, data []byte, fun func(*os.File)) {
	f, err := ioutil.TempFile("", pat)
	if err != nil {
		t.Error()
		return
	}

	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		t.Error()
		return
	}

	fun(f)
}

func TestDefaultValueHandler(t *testing.T) {
	data := []byte("this is some data")
	edata := []byte("\"this is some data\"")
	withTestFile(t, "*", data, func(f *os.File) {
		rdata, err := DefaultValueHandler(f.Name())
		if err != nil {
			t.Error()
			return
		}

		if !bytes.Equal(rdata, edata) {
			t.Error()
		}
	})

	d, err := DefaultValueHandler("skdflksdfjkdsfkldsjfkkfnkn")
	if d != nil || err == nil {
		t.Error()
	}
}

func TestStripExtKeyHandler(t *testing.T) {
	if fn := StripExtKeyHandler("test.json"); fn != "test" {
		t.Error()
	}

	if fn := StripExtKeyHandler("test"); fn != "test" {
		t.Error()
	}

	if fn := StripExtKeyHandler(".json"); fn != "" {
		t.Error()
	}

	if fn := StripExtKeyHandler(""); fn != "" {
		t.Error()
	}
}

func TestJsonValueHandler(t *testing.T) {
	data := []byte("{\"a\": 2}")
	withTestFile(t, "*.json", data, func(f *os.File) {
		rdata, err := JsonValueHandler(f.Name())
		if err != nil {
			t.Error()
			return
		}

		if !bytes.Equal(rdata, data) {
			t.Error()
		}
	})

	d, err := JsonValueHandler("skdflksdfjkdsfkldsjfkkfnkn")
	if d != nil || err == nil {
		t.Error()
	}
}
