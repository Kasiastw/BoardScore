package main

import (

	"github.com/labstack/echo"
	"net/http"
)

type message struct {
	greetings string `json:"greetings"`
}

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})

	e.POST("api/v1/greetings", func(ctx echo.Context) error {

		m := &message{greetings: "Hello!"}

		return ctx.JSON(http.StatusOK, m)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
