package main

import (
	"greedy-game-test/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.Set(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
