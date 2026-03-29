package signalling

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	RoomID string
	Conn   *websocket.Conn
}
type Hub struct {
	Rooms map[string]map[string]*Client
	Mu    sync.Mutex
}

var HubInstance = &Hub{
	Rooms: make(map[string]map[string]*Client),
}
