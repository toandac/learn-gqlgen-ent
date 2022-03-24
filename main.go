package main

import (
	"context"
	"errors"
	"gqlgen-ent/ent"
	"gqlgen-ent/ent/artcile"
	"gqlgen-ent/ent/user"
	"gqlgen-ent/graph"
	"log"
	"net/http"

	handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "toannd",
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "3306",
		DBName:               "gqlgen-ent",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("Error: mysql client: %v\n", err)
	}
	defer client.Close()
	// Run the migration here
	if err := client.Schema.Create(context.Background()); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	// Configure the GraphQL server and start
	srv := handler.NewDefaultServer(graph.NewSchema(client))
	{
		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	e.POST("/users", func(c echo.Context) error {
		// Create an article entity
		a, err := client.Artcile.Create().
			SetTitle("title 1").
			SetDescription("description 1").
			Save(c.Request().Context())
		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed creating article %v\n", err)
		}

		u, err := client.User.Create().SetName("toannd").SetAge(26).AddArticles(a).Save(c.Request().Context())
		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed creating user %v\n", err)
		}

		return c.JSON(http.StatusCreated, u)
	})

	e.GET("/users", func(c echo.Context) error {
		u, err := client.User.
			Query().
			WithArticles(). // WithArticles loads articles
			Where(user.IDEQ(2)).
			Only(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed quering users %v\n", err)
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/article/user", func(c echo.Context) error {
		a, err := client.Artcile.
			Query().
			Where(artcile.IDEQ(1)).
			QueryUser(). // Query user
			Only(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed quering user %v\n", err)
		}

		return c.JSON(http.StatusOK, a)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
