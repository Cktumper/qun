package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"peon.top/qun/config"
	"peon.top/qun/controllers/chat"
	"peon.top/qun/controllers/system"
	"strings"
)

type Http struct {
	engine *gin.Engine
}

func NewHTTP() *Http {
	return &Http{
		engine: gin.Default(),
	}
}

//	获取服务实例
func (p *Http) GetService() *gin.Engine {
	return p.engine
}

//	加载 HTTP 路由
func (p *Http) LoadRouter() *Http {
	p.engine.LoadHTMLFiles(fmt.Sprintf("%s../resources/chat/index.html", p.getCurrentDirectory()))

	//	加载系统路由
	p.loadSystemRouter()

	p.engine.GET("/chat", chat.NewIndexController().Index)
	p.engine.GET("/chat/login", chat.NewLoginController().Login)

	return p
}

//	加载系统路由
func (p *Http) loadSystemRouter() {
	router := p.engine.Group("info")

	router.GET("version", system.NewInformationController().Version)
	router.GET("ping", system.NewInformationController().Ping)
}

//	监听 HTTP 服务
func (p *Http) Listen() error {
	//return endless.ListenAndServe(fmt.Sprintf("%s:%s", config.ReadString("SERVICE_HOST"), config.ReadString("SERVICE_PORT")), p.engine)
	return p.engine.Run(fmt.Sprintf("%s:%s", config.ReadString("SERVICE_HOST"), config.ReadString("SERVICE_PORT")))
}

//	加载当前执行的目录
func (p *Http) getCurrentDirectory() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return path[0 : i+1]
}

func Unescaped(str string) template.HTML {
	return template.HTML(str)
}

//	Start
func (p *Http) Start() error {
	p.LoadRouter()
	return p.Listen()
}

//	Close
func (p *Http) End() error {
	return nil
}
