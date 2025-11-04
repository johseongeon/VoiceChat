package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

// Helper to make Gorilla Websockets threadsafe.
type ThreadSafeWriter struct {
	*websocket.Conn
	sync.Mutex
}

func (t *ThreadSafeWriter) WriteJSON(v any) error {
	t.Lock()
	defer t.Unlock()

	return t.Conn.WriteJSON(v)
}
