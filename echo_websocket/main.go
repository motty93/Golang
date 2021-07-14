package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	clients   = sync.Map{}
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Message struct {
	Text   string `json:"text"`
	Type   int    `json:"type"`
	Action string `json:"action"`
}

func websocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	// clientを登録
	clients.Store(ws, true)

	for {
		var msg Message
		// websocketを介して取得したメッセージの読み取り
		if err := ws.ReadJSON(&msg); err != nil {
			clients.Delete(ws)
			c.Logger().Error(err)
		}
		c.Logger().Printf("%s sent: %s", ws.RemoteAddr(), msg.Text)

		broadcast <- msg
	}
}

func messagesHandler() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		msg := <-broadcast
		// 接続しているクライアント全てにメッセージの送信
		clients.Range(func(k, v interface{}) bool {
			// interface型から*websocket.Connへアサーション
			client := k.(*websocket.Conn)
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				clients.Delete(client)
				return false
			}
			return true
		})
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
