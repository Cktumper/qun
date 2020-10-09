package main

import (
	"peon.top/qun/bootstrap"
	"peon.top/qun/config"
	"peon.top/qun/devices/http"
	"peon.top/qun/devices/message"
	"peon.top/qun/devices/session"
)

var bp = bootstrap.New()

//	main function
//	You Know
func main() {
	//	Load Env
	config.NewEnv(".env").Load()

	//	Load Devices
	bp.Add(http.NewHTTP()).
		Add(session.NewService()).
		Add(message.NewService())

	//	Run and waiting
	bp.Run().Waiting()
}
