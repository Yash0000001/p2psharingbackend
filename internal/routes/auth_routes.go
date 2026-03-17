package routes

import (
	"net/http"

	"github.com/yash0000001/p2psharingbackend/internal/controller"
	"github.com/yash0000001/p2psharingbackend/internal/middleware"
	"github.com/yash0000001/p2psharingbackend/internal/utils"
)

func AuthRoutes() {
	http.Handle(
		"/auth/signup",
		middleware.RateLimit(http.HandlerFunc(utils.POSTOnly(controller.Signup))),
	)

	http.Handle(
		"/auth/login",
		middleware.RateLimit(http.HandlerFunc(utils.POSTOnly(controller.Login))),
	)

	http.Handle(
		"/auth/google",
		middleware.RateLimit(http.HandlerFunc(utils.POSTOnly(controller.GoogleSignin))),
	)

	http.Handle(
		"/auth/reset-password",
		middleware.RateLimit(http.HandlerFunc(utils.POSTOnly(controller.ResetPassword))),
	)
	http.Handle(
		"/auth/forgot-password",
		middleware.RateLimit(http.HandlerFunc(utils.POSTOnly(controller.ForgotPassword))),
	)
}
