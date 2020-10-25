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

//	1. 接下来的工作，实现好看的前端页面上线，单ROOM执行
//	2. 实现 MYSQL 管理 ROOM 与 用户系统，增加 HTTP SESSION 系统
//		连接使用 SESSION 而不是随便打昵称
//	3. 实现用户注册、登陆功能
//	4. 实现机器人接口-更新核心功能
//	5. 实现 HTTP 机器人
//	6. 实现多 ROOM 功能，可支持让用户申请并开启 ROOM 功能
//	7. 实现持久化功能 ROOM, USER ,CONNECTION
