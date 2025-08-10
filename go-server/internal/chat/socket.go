package chat

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConnectToClient(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	ReadMessageFromConnections(conn)
	return nil
}

func WriteMessageToConnection(conn *websocket.Conn, messageType int ,message []byte) error {
	err := conn.WriteMessage(messageType, message)
	if err != nil {
		fmt.Println("error encountered", err)
	}
	return nil
}

func ReadMessageFromConnections(conn *websocket.Conn) error {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Caught error while reading the message", err)
			return err
		}

		err = WriteMessageToConnection(conn, msgType, msg)
		if err != nil {
			fmt.Println("error found", err)
			return err
		}
	}

}
