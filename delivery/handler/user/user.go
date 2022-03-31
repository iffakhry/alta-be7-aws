package user

import (
	"be7/layered/delivery/helper"
	_middlewares "be7/layered/delivery/middlewares"
	_entities "be7/layered/entities"
	_userUseCase "be7/layered/usecase/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(userUseCase _userUseCase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all users", users))
	}
}

func (uh *UserHandler) GetByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		// check apakah id dari token sama dengan id dari param
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		users, rows, err := uh.userUseCase.GetById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all users", users))
	}
}

func (uh *UserHandler) GetHelloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success get hello feature"))
	}
}

func (uh *UserHandler) PostInserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		errBind := c.Bind(&user)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed bind data"))
		}
		user, row, err := uh.userUseCase.Insert(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create user"))
		}
		if row == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create user"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create users"))
	}
}
