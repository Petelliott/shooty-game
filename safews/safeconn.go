package safews

import (
	"github.com/gorilla/websocket"
	"sync"
	"net"
)

type SafeConn struct {
    conn *websocket.Conn
    mu sync.Mutex
}

func MakeSafe(conn *websocket.Conn) *SafeConn {
	var sc SafeConn
	sc.conn = conn
	return &sc
}

func (c *SafeConn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.Close()
}

func (c *SafeConn) LocalAddr() net.Addr {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.LocalAddr()
}

func (c *SafeConn) ReadJSON(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.ReadJSON(v)
}

func (c *SafeConn) ReadMessage() (messageType int, p []byte, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.ReadMessage()
}

func (c *SafeConn) RemoteAddr() net.Addr {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.RemoteAddr()
}

func (c *SafeConn) WriteJSON(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.WriteJSON(v)
}

func (c *SafeConn) WriteMessage(messageType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.WriteMessage(messageType, data)
}
