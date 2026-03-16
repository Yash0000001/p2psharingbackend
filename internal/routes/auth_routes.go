package routes

import (
	"net/http"

	"github.com/yash0000001/p2psharingbackend/internal/controller"
	"github.com/yash0000001/p2psharingbackend/internal/middleware"
)

func AuthRoutes() {
	http.Handle(
		"/auth/signup",
		middleware.RateLimit(http.HandlerFunc(controller.Signup)),
	)

	http.Handle(
		"/auth/login",
		middleware.RateLimit(http.HandlerFunc(controller.Login)),
	)

	http.Handle(
		"/auth/google",
		middleware.RateLimit(http.HandlerFunc(controller.GoogleSignin)),
	)
}
