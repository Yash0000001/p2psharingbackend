package signalling

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("roomId")
	userID := r.URL.Query().Get("userId")

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}
	client := &Client{
		ID:     userID,
		RoomID: roomID,
		Conn:   conn,
	}

	registerClient(client)
	go handleMessages(client)
}
func registerClient(c *Client) {
	HubInstance.Mu.Lock()
	defer HubInstance.Mu.Unlock()

	if HubInstance.Rooms[c.RoomID] == nil {
		HubInstance.Rooms[c.RoomID] = make(map[string]*Client)
	}

	// Notify existing users
	for _, client := range HubInstance.Rooms[c.RoomID] {
		client.Conn.WriteJSON(map[string]interface{}{
			"type":   "user-joined",
			"userId": c.ID,
		})
	}

	HubInstance.Rooms[c.RoomID][c.ID] = c
}
func handleMessages(c *Client) {
	defer func() {
		removeClient(c)
		c.Conn.Close()
	}()

	for {
		var msg map[string]interface{}
		if err := c.Conn.ReadJSON(&msg); err != nil {
			break
		}

		targetID := msg["target"].(string)

		HubInstance.Mu.Lock()
		targetClient := HubInstance.Rooms[c.RoomID][targetID]
		HubInstance.Mu.Unlock()

		if targetClient != nil {
			targetClient.Conn.WriteJSON(msg)
		}
	}
}
func removeClient(c *Client) {
	HubInstance.Mu.Lock()
	defer HubInstance.Mu.Unlock()

	delete(HubInstance.Rooms[c.RoomID], c.ID)

	for _, client := range HubInstance.Rooms[c.RoomID] {
		client.Conn.WriteJSON(map[string]interface{}{
			"type":   "user-left",
			"userId": c.ID,
		})
	}
}
