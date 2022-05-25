package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controllers "github.com/oxwazz/twitter/app/deliveries"
	"github.com/oxwazz/twitter/app/entities/databases"
	"github.com/oxwazz/twitter/db"
	"github.com/oxwazz/twitter/graph"
	"github.com/oxwazz/twitter/graph/generated"
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
