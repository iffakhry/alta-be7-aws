package auth

import (
	"be7/layered/delivery/helper"
	_authUseCase "be7/layered/usecase/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUseCase _authUseCase.AuthUseCaseInterface
}

func NewAuthHandler(auth _authUseCase.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: auth,
	}
}

func (ah *AuthHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// var login _entities.User
		type loginData struct {
			Identifier string `json:"identifier"`
			Password   string `json:"password"`
		}
		var login loginData
		// err := c.Bind(&login)
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error bind data"))
		}
		token, errorLogin := ah.authUseCase.Login(login.Identifier, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success login", responseToken))
	}
}
