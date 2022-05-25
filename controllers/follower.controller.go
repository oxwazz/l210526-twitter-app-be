package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/oxwazz/l210526-twitter-app-be/models"

	"github.com/labstack/echo"
)

func FetchAllFollower(c echo.Context) error {
	result, err := models.FetchAllFollower()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreFollower(c echo.Context) error {
	follower := new(models.Follower)

	fmt.Println("sdfsdf")
	if err := c.Bind(follower); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	fmt.Println(follower)

	result, err := models.StoreFollower(
		follower.FollowedId,
		follower.FollowerId,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateFollower(c echo.Context) error {
	follower := new(models.Follower)

	fmt.Println("sdfsdf")
	if err := c.Bind(follower); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	fmt.Println(follower)

	conv_id, err := strconv.Atoi(follower.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateFollower(
		conv_id,
		follower.FollowedId,
		follower.FollowerId,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteFollower(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteFollower(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
