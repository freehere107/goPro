package routes

import "chat/models"

type MRouter struct {
	ChatConn map[string]ChatConn
}

type ChatConn struct {
	Conn chan models.Message
}

func NewMessageRouter() *MRouter {
	r := &MRouter{ChatConn: make(map[string]ChatConn)}
	appName := []string{"chat"}
	go routerSub(r, appName)
	return r
}

//频道分发
func routerSub(r *MRouter, channel []string) {

}
