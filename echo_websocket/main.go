package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{}
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func websocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// 初回メッセージ
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	if err != nil {
		c.Logger().Error(err)
	}

	for {
		// websocketを介して取得したメッセージの読み取り
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

		// 受け取ったメッセージをクライアント側へ表示
		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			c.Logger().Error(err)
		}
		// ここから続き
		// Messageの構造体へmsgを入れbroadcastへ渡す
		broadcast <- string(msg)
	}
}

func handleMessages() {
	for {

	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", websocketHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
