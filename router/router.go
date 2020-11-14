package router

import (
	"greedy-game-test/handlers"

	"github.com/labstack/echo/v4"
)

//Set sets all routes
func Set(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/insert", handlers.Insert)
	v1.GET("/query", handlers.Query)
}
