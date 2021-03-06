package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	controllers "github.com/oxwazz/l210526-twitter-app-be/app/deliveries"
	"github.com/oxwazz/l210526-twitter-app-be/app/entities/databases"
	"github.com/oxwazz/l210526-twitter-app-be/db"
	"github.com/oxwazz/l210526-twitter-app-be/graph"
	"github.com/oxwazz/l210526-twitter-app-be/graph/generated"
	"github.com/oxwazz/l210526-twitter-app-be/helpers"
)

func main() {

	port := helpers.GetEnv("PORT")

	databases.Init()

	db.Init()

	e := controllers.Init()

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":" + port))

}
