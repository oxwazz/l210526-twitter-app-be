package controllers

import (
	"github.com/labstack/echo/middleware"
	controllers2 "github.com/oxwazz/twitter/app/controllers"
	"net/http"

	"github.com/oxwazz/twitter/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello :)")
	})

	// TODO: register
	e.GET("/register", controllers2.Register)

	// TODO: login f
	e.POST("/login", controllers2.Login)

	e.GET("/users", controllers2.FetchAllUser, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret-key"),
	}))
	e.PUT("/api/v1/users", controllers.UpdateUser)
	e.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	e.GET("/api/v1/tweets", controllers.FetchAllTweet)
	e.POST("/api/v1/tweets", controllers.StoreTweet)
	e.PUT("/api/v1/tweets", controllers.UpdateTweet)
	e.DELETE("/api/v1/tweets/:id", controllers.DeleteTweet)

	e.GET("/api/v1/followers", controllers.FetchAllFollower)
	e.POST("/api/v1/followers", controllers.StoreFollower)
	e.PUT("/api/v1/followers", controllers.UpdateFollower)
	e.DELETE("/api/v1/followers/:id", controllers.DeleteFollower)

	return e
}
