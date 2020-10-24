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

//	接下来工作
//	实现聊天程序，使用 WEBSOCKET ，实现应用层聊天系统
//	实现 Session ,Packet ,Connection ,User 四大接口
