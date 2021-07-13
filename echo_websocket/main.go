package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Message struct {
	Text string `json:"text"`
	Type int    `json:"type"`
}

func websocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	// clientを登録
	clients[ws] = true

	for {
		// websocketを介して取得したメッセージの読み取り
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			c.Logger().Error(err)
		}
		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

		broadcast <- Message{Text: string(msg), Type: msgType}
	}
}

func messagesHandler() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		msg := <-broadcast
		// 接続しているクライアント全てにメッセージの送信
		for client := range clients {
			if err := client.WriteMessage(msg.Type, []byte(msg.Text)); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
				break
			}
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/ws", websocketHandler)
	go messagesHandler()
	e.Logger.Fatal(e.Start(":8080"))
}
