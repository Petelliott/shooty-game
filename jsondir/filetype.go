package jsondir

import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type keyhandler func(string) string
type valuehandler func(string) ([]byte, error)

type FileEncoder struct {
	keyhandlers map[string]keyhandler
	defaultKeyHandler keyhandler

	valuehandlers map[string]valuehandler
	defaultValueHandler valuehandler
}

func NewEncoder() FileEncoder {
	return FileEncoder{
		make(map[string]keyhandler),
		defaultKeyHandler,
		make(map[string]valuehandler),
		defaultValueHandler,
	}
}

func (fe *FileEncoder) AddKeyHandler(ext string, handler keyhandler) {
	fe.keyhandlers[ext] = handler
}

func (fe *FileEncoder) AddValueHandler(ext string, handler valuehandler) {
	fe.valuehandlers[ext] = handler
}

func (fe *FileEncoder) SetDefaultKeyHandler(handler keyhandler) {
	fe.defaultKeyHandler = handler
}

func (fe *FileEncoder) SetDefaultValueHandler(handler valuehandler) {
	fe.defaultValueHandler = handler
}

func (fe *FileEncoder) GetKey(path string) string {
	 ext := filepath.Ext(path)
	 if handler, ok := fe.keyhandlers[ext]; ok {
		return handler(path)
	 } else {
		return fe.defaultKeyHandler(path)
	 }
}

func (fe *FileEncoder) Encode(path string) ([]byte, error) {
	 ext := filepath.Ext(path)
	 if handler, ok := fe.valuehandlers[ext]; ok {
		return handler(path)
	 } else {
		return fe.defaultValueHandler(path)
	 }
}

func DefaultEncoder() FileEncoder {
	fe := NewEncoder()
	fe.AddKeyHandler(".json", stripExtKeyHandler)
	fe.AddValueHandler(".json", jsonValueHandler)
	return fe
}

func defaultKeyHandler(filename string) string {
	return filename
}

func defaultValueHandler(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(data))
}

func stripExtKeyHandler(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func jsonValueHandler(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

