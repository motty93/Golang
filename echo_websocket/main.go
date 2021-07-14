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
	broadcast = make(chan User)
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type User struct {
	Message string          `json:"message"`
	Action  string          `json:"action"`
	Name    string          `json:"name"`
	Conn    *websocket.Conn `json:"-"`
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
		var user User
		// websocketを介して取得したメッセージの読み取り
		if err := ws.ReadJSON(&user); err != nil {
			clients.Delete(ws)
			c.Logger().Error(err)
		}
		c.Logger().Printf("%s sent: %s", ws.RemoteAddr(), user.Message)

		user.Conn = ws
		broadcast <- user
	}
}

func messagesHandler() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		data := <-broadcast
		// 接続しているクライアント全てにメッセージの送信
		switch data.Action {
		case "broadcast":
			clients.Range(func(k, v interface{}) bool {
				// interface型から*websocket.Connへアサーション
				client := k.(*websocket.Conn)
				if err := client.WriteJSON(data); err != nil {
					log.Printf("error: %v", err)
					client.Close()
					clients.Delete(client)
					return false
				}
				return true
			})
		case "left":
			// clientsからconnection ユーザーを削除
			clients.Delete(data.Conn)
			data.Conn.Close()
		case "users":
			log.Println("all users list")
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
