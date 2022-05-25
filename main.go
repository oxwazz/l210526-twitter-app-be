package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controllers "github.com/oxwazz/l210526-twitter-app-be/app/deliveries"
	"github.com/oxwazz/l210526-twitter-app-be/app/entities/databases"
	"github.com/oxwazz/l210526-twitter-app-be/db"
	"github.com/oxwazz/l210526-twitter-app-be/graph"
	"github.com/oxwazz/l210526-twitter-app-be/graph/generated"
	"github.com/spf13/viper"
)

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	valueport, ok := viper.Get("PORT").(string)
	if !ok {
		fmt.Println("error ges")
	}

	//port := os.Getenv("PORT")

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

	e.Logger.Fatal(e.Start(":" + valueport))

}
