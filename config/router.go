package config

import (
	"github.com/dharmasastra/realtime-chat-go/app/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo{
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} | ${error}\n",
	}))
	e.Use(middleware.Recover())

	// Configure websocket route
	e.GET("/ws", controllers.HandleConnections)

	return e
}
