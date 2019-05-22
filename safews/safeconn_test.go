package safews

import (
	"testing"
	"github.com/gorilla/websocket"
	"context"
	"net/http"
	"fmt"
	"strings"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func connPair(t *testing.T) (*SafeConn, *SafeConn, func()) {
	var conns *websocket.Conn
	var connc *websocket.Conn

	srv := &http.Server{Addr: "127.0.0.1:8080"}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var err error
		conns, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			t.Error()
			return
		}

		for {}
	})

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			t.Error()
		}
	}()

	connc, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
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
	defer clean()

	if sc == nil || cc == nil {
		t.Error()
		return
	}

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
