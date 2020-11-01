package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Service struct {
	//	公有方法
	EnableTLS bool

	//	私有方法
	engine *gin.Engine
	route  *Route
	host   string
	port   int
}

//	构建一个WEB Socket服务
//
//	Author(Wind)
func NewService(host string, port int) *Service {
	//	构建服务对象
	p := &Service{
		engine: gin.Default(),
		port:   port,
		host:   host,
	}

	//	加载路由对象
	p.route = NewRoute(p.engine)

	//	返回自己
	return p
}

//	开启服务
//
//	Author(Wind)
func (p *Service) Start() error {
	p.engine.LoadHTMLGlob(fmt.Sprintf("%s../templates/*.html", p.getCurrentDirectory()))
	p.engine.Static("static", fmt.Sprintf("%s../templates/", p.getCurrentDirectory()))

	//	加载路由表
	p.route.Load()

	//	判定是否需要使用 HTTPS
	if p.EnableTLS {
		return ErrNotSupportTLS
	}

	//	使用 HTTP 监听端口
	return p.engine.Run(fmt.Sprintf("%s:%d", p.host, p.port))
}

//	关闭服务
//
//	Author(Wind)
func (p *Service) End() error {
	return nil
}

//	加载当前执行的目录
func (p *Service) getCurrentDirectory() string {
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
