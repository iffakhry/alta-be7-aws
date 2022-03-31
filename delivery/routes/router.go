package routes

import (
	_authHandler "be7/layered/delivery/handler/auth"
	_userHandler "be7/layered/delivery/handler/user"
	_middlewares "be7/layered/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.GET("/hello", uh.GetHelloHandler())
	e.GET("/users", uh.GetAllHandler())
	e.GET("/users/:id", uh.GetByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.PostInserHandler())
}
