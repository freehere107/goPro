package main

import (
	"chat/models"
	"chat/routes"
	"chat/utiles"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

var (
	upGrader = websocket.Upgrader{}
)

func ReceiveMessage(c *websocket.Conn) {
	redisCoon := utiles.RedisClient
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		var m models.Message
		if err := json.Unmarshal([]byte(msg), &m); err != nil {
			errMsg := fmt.Sprintf("Dispatch error:%s", err)
			errors.New(errMsg)
		}
		if err == nil {
			redisCoon.Publish("chat", m.Content).Result()
		}
	}
}

func SocketReceive(messageRoute *routes.MRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.FormValue("user_id")
		chatChannel := make(chan models.Message)
		messageRoute.ChatConn[userID] = routes.ChatConn{Conn: chatChannel}
		c, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		ReceiveMessage(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	messageRoute := routes.NewMessageRouter()
	e.GET("/ws", standard.WrapHandler(http.HandlerFunc(SocketReceive(messageRoute))))
	e.Run(standard.New(":1323"))
}
