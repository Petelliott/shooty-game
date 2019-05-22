package safews

import (
	"testing"
	"github.com/gorilla/websocket"
	"context"
	"net/http"
	"fmt"
	"strings"
	"time"
	"math/rand"
	"strconv"
	"bytes"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func connPair(t *testing.T) (*SafeConn, *SafeConn, func()) {
	var conns *websocket.Conn
	var connc *websocket.Conn

	srv := &http.Server{Addr: "0.0.0.0:8080"}

	slug := "/" + strconv.Itoa(rand.Int())

	http.HandleFunc(slug, func(w http.ResponseWriter, r *http.Request) {
		var err error
		conns, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			t.Error()
			return
		}
	})

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			t.Error()
		}
	}()

	time.Sleep(1 * time.Second)
	connc, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080"+slug, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil, func() {
			conns.Close()
			connc.Close()
			srv.Shutdown(context.TODO())
		}
	}

	sconns := MakeSafe(conns)
	sconnc := MakeSafe(connc)

	return sconns, sconnc, func() {
		sconnc.Close()
		sconns.Close()
		srv.Shutdown(context.TODO())
	}
}

func TestLocalAddr(t *testing.T) {
	sc, cc, clean := connPair(t)
	if sc == nil || cc == nil {
		t.Error()
		return
	}

	defer clean()

	st := sc.LocalAddr().String()
	if  st != "127.0.0.1:8080" && st != "[::1]:8080" {
		fmt.Println(sc.LocalAddr().String())
		t.Error()
	}

	ct := cc.LocalAddr().String()
	if !strings.HasPrefix(ct, "127.0.0.1") && !strings.HasPrefix(ct, "[::1]") {
		fmt.Println(cc.LocalAddr().String())
		t.Error()
	}
}

func TestRemoteAddr(t *testing.T) {
	sc, cc, clean := connPair(t)
	if sc == nil || cc == nil {
		t.Error()
		return
	}

	defer clean()

	ct := cc.RemoteAddr().String()
	if  ct != "127.0.0.1:8080" && ct != "[::1]:8080" {
		fmt.Println(ct)
		t.Error()
	}

	st := sc.RemoteAddr().String()
	if !strings.HasPrefix(st, "127.0.0.1") && !strings.HasPrefix(st, "[::1]") {
		fmt.Println(st)
		t.Error()
	}
}

type thing struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func TestReadWriteJson(t *testing.T) {
	sc, cc, clean := connPair(t)
	if sc == nil || cc == nil {
		t.Error()
		return
	}

	defer clean()

	th := thing{5, 7}
	err := sc.WriteJSON(th)
	if err != nil {
		t.Error()
		return
	}

	var t2 thing
	err = cc.ReadJSON(&t2)
	if err != nil {
		t.Error()
		return
	}

	if t2.X != 5 || t2.Y != 7 {
		t.Error()
	}
}

func TestReadWriteMessage(t *testing.T) {
	sc, cc, clean := connPair(t)
	if sc == nil || cc == nil {
		t.Error()
		return
	}

	defer clean()

	b1 := []byte("hello world\n")
	err := cc.WriteMessage(websocket.BinaryMessage, b1)
	if err != nil {
		t.Error()
		return
	}

	mt, t2, err := sc.ReadMessage()
	if err != nil {
		t.Error()
		return
	}

	if mt != websocket.BinaryMessage {
		t.Error()
	}

	if !bytes.Equal(t2, b1) {
		t.Error()
	}
}
