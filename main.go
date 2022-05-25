package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controllers "l210526-twitter-app-be/app/deliveries"
	"l210526-twitter-app-be/app/entities/databases"
	"l210526-twitter-app-be/db"
	"l210526-twitter-app-be/graph"
	"l210526-twitter-app-be/graph/generated"
)

func main() {

	databases.Init()

	db.Init()

	e := controllers.Init()

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		fmt.Println(333310, "sdfsdf")
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	}, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret-key"),
	}))

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":1234"))

}
