package chat

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type Client struct {
	ID      string
	EchoCtx *echo.Context
	Conn    *websocket.Conn
	Pool    *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Read
		msg := ""
		err := websocket.Message.Receive(c.Conn, &msg)
		if err != nil {
			(*(c.EchoCtx)).Logger().Error(err)
			break
		}
		fmt.Printf("%s\n", msg)
	}
}

func (c *Client) Send(message Message) {

}
