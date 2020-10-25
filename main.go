package main

import (
	"peon.top/qun/bootstrap"
	"peon.top/qun/chat"
	"peon.top/qun/config"
	"peon.top/qun/services/websocket"
)

var bp = bootstrap.New()

//	main function
//	You Know
func main() {
	//	Load Env
	config.NewEnv(".env").Load()

	bp.Add(chat.NewService()).
		Add(websocket.NewService(8080))

	//	Run and waiting
	bp.Run().Waiting()
}
