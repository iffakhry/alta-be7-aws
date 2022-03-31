package main

import (
	"be7/layered/configs"
	_authHandler "be7/layered/delivery/handler/auth"
	_userHandler "be7/layered/delivery/handler/user"
	_authRepository "be7/layered/repository/auth"
	_userRepository "be7/layered/repository/user"
	_authUseCase "be7/layered/usecase/auth"
	_userUseCase "be7/layered/usecase/user"
	"fmt"
	"log"

	_middlewares "be7/layered/delivery/middlewares"
	_routes "be7/layered/delivery/routes"
	_utils "be7/layered/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())

	_routes.RegisterPath(e, userHandler)
	_routes.RegisterAuthPath(e, authHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
