package main

import (
	"github.com/GUEThe/go-and-vue3-webchat-backend/pkg/chat"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	pool := chat.NewPool()
	pool.init()
	e.GET("/ws", func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			client := chat.Client{ID: "sss", EchoCtx: &c, Conn: ws, Pool: pool}
			client.Read()
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
