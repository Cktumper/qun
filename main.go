package main

import (
	"peon.top/qun/bootstrap"
	"peon.top/qun/config"
)

var bp = bootstrap.New()

//	main function
//	You Know
func main() {
	//	Load Env
	config.NewEnv(".env").Load()

	//	Run and waiting
	bp.Run().Waiting()
}
