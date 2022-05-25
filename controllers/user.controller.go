package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/oxwazz/l210526-twitter-app-be/models"

	"github.com/labstack/echo"
)

func UpdateUser(c echo.Context) error {
	user := new(models.User)

	fmt.Println("sdfsdf")
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	fmt.Println(user)

	return c.JSON(http.StatusOK, "result")
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
