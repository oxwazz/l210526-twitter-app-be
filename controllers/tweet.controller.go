package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"l210526-twitter-app-be/models"

	"github.com/labstack/echo"
)

func FetchAllTweet(c echo.Context) error {
	result, err := models.FetchAllTweet()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTweet(c echo.Context) error {
	tweet := new(models.Tweet)

	if err := c.Bind(tweet); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}
	fmt.Println(tweet)
	conv_id, err := strconv.Atoi(tweet.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.StoreTweet(
		conv_id,
		tweet.Content,
		tweet.Attachment,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTweet(c echo.Context) error {
	tweet := new(models.Tweet)

	fmt.Println("sdfsdf")
	if err := c.Bind(tweet); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": err.Error()})
	}

	conv_id, err := strconv.Atoi(tweet.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateTweet(
		conv_id,
		tweet.Content,
		tweet.Attachment,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTweet(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteTweet(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
