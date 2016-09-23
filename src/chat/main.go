package main

import (
	"chat/utiles"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}
)

func ReceiveMessage(c *websocket.Conn) {
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func SocketReceive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.FormValue("user_id")
		fmt.Println(user_id)
		c, err := upgrader.Upgrade(w, r, nil)
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
	result, _ := utiles.RedisClient.Get("key").Result()
	println(result)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", standard.WrapHandler(http.HandlerFunc(SocketReceive())))
	e.Run(standard.New(":1323"))
}
