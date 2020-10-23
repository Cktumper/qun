package main

import (
	"peon.top/qun/bootstrap"
	"peon.top/qun/config"
	"peon.top/qun/devices/example"
)

var bp = bootstrap.New()

//	main function
//	You Know
func main() {
	//	Load Env
	config.NewEnv(".env").Load()

	bp.Add(example.NewExample())

	//	Run and waiting
	bp.Run().Waiting()
}
