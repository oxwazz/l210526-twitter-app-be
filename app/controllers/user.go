package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/oxwazz/twitter/app/entities"
	"github.com/oxwazz/twitter/app/usecases"
	"net/http"
)

func FetchAllUser(c echo.Context) error {
	var res Response
	users, err := usecases.FetchAllUser()
	if err != nil {
		res.Success = false
		res.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Success = true
	res.Data = users
	return c.JSON(http.StatusOK, res)

}

func Register(c echo.Context) error {
	var res Response

	user := new(entities.User)

	if err := c.Bind(user); err != nil {
		fmt.Println(err)
		res.Message = err.Error()
		return c.JSON(http.StatusOK, res)
	}

	users, err := usecases.Register(*user)
	if err != nil {
		res.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Success = true
	res.Data = &users
	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	var res Response

	reqBody := entities.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		res.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}

	token, err := usecases.Login(reqBody.Email, reqBody.Username, reqBody.Phone, reqBody.Password)
	if err != nil {
		res.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Success = true
	res.Data = map[string]interface{}{
		"access_token": token,
	}
	return c.JSON(http.StatusOK, res)
	//
	//user := new(entities.User)
	//
	//if err := c.Bind(user); err != nil {
	//	fmt.Println(err)
	//	res.Message = err.Error()
	//	return c.JSON(http.StatusOK, res)
	//}
	//
	//users, err := usecases.Register(*user)
	//if err != nil {
	//	res.Message = err.Error()
	//	return c.JSON(http.StatusInternalServerError, res)
	//}
	//
	//res.Success = true
	//res.Data = &users
	//return c.JSON(http.StatusOK, res)
}
