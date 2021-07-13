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
	clients.Store(ws, true)

	for {
		// websocketを介して取得したメッセージの読み取り
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			clients.Delete(ws)
			c.Logger().Error(err)
		}
		c.Logger().Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

		broadcast <- Message{Text: string(msg), Type: msgType}
	}
}

func messagesHandler() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		msg := <-broadcast
		// 接続しているクライアント全てにメッセージの送信
		clients.Range(func(client, value interface{}) bool {
			if err := client.WriteMessage(msg.Type, []byte(msg.Text)); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				clients.Delete(client)
				return false
			}
			return true
		})
		// for client := range clients {
		// 	if err := client.WriteMessage(msg.Type, []byte(msg.Text)); err != nil {
		// 		log.Printf("error: %v", err)
		// 		client.Close()
		// 		delete(clients, client)
		// 	}
		// }
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
