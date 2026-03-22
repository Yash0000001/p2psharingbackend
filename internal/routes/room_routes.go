package routes

import (
	"net/http"

	"github.com/yash0000001/p2psharingbackend/internal/controller"
	"github.com/yash0000001/p2psharingbackend/internal/middleware"
	"github.com/yash0000001/p2psharingbackend/internal/utils"
)

func RoomRoutes() {

	http.Handle(
		"/room/create",
		middleware.RateLimit(
			middleware.JWTAuth(
				utils.POSTOnly(controller.CreateRoom),
			),
		),
	)

	http.Handle(
		"/room/join",
		middleware.RateLimit(
			middleware.JWTAuth(
				utils.POSTOnly(controller.JoinRoom),
			),
		),
	)

	http.Handle(
		"/room/delete",
		middleware.RateLimit(
			middleware.JWTAuth(
				utils.DELETEOnly(controller.DeleteRoom),
			),
		),
	)

	http.Handle(
		"/room/nearby",
		middleware.RateLimit(
			middleware.JWTAuth(
				utils.GETOnly(controller.GetNearbyRooms),
			),
		),
	)

	http.Handle(
		"/room/leave",
		middleware.RateLimit(
			middleware.JWTAuth(
				utils.POSTOnly(controller.LeaveRoom),
			),
		),
	)
}
