package main

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func handleMsg(ws *websocket.Conn, c echo.Context) {
	for {
		// Write
		err := websocket.Message.Send(ws, "Hello, Client!")
		if err != nil {
			c.Logger().Error(err)
			break
		}

		// Read
		msg := ""
		err = websocket.Message.Receive(ws, &msg)
		if err != nil {
			c.Logger().Info(err == io.EOF)
			c.Logger().Error(err)
			break
		}
		fmt.Printf("%s\n", msg)
	}
}

func wsHandler(c echo.Context) error {
	// userWsMap := make(map[string]*websocket.Conn)
	websocket.Handler(func(ws *websocket.Conn) {
		// client := chat.Client{}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", wsHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
