package routes

import (
	"net/http"

	"github.com/yash0000001/p2psharingbackend/internal/signalling"
)

func SignallingRoutes() {
	http.HandleFunc("/ws", signalling.HandleWebsocket)
}
