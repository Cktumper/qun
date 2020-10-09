package bootstrap

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type Boost struct {
	*Config
	*Command

	boots []Bootstrap
	stop  chan os.Signal
}

func New() *Boost {
	return &Boost{
		Command: NewCommand(),
		Config:  NewConfig(),
	}
}

//	添加启动设备
//
//	Author(Wind)
func (p *Boost) Add(b Bootstrap) *Boost {
	p.boots = append(p.boots, b)
	return p
}

//	启动设备
//
//	Author(Wind)
func (p *Boost) Run() *Boost {
	//	遍历各个协程逐个进行启动
	for _, boot := range p.boots {
		go boot.Start()
	}

	//	判定如果是命令模式则执行命令
	if p.CommandMode() {
		//	延迟一秒，等待服务启动
		time.Sleep(1 * time.Second)

		//	执行命令
		for _, cmd := range p.GetShouldExecCommands() {
			if err := cmd.Handle(); err != nil {
				log.Fatalln(err)
			}
		}
	}

	return p
}

//	等待关闭信号
//
//	Author(Wind)
func (p *Boost) Waiting() {
	//	如果是命令模式则不等待
	if p.CommandMode() {
		return
	}

	//	设置监听信号量
	p.stop = make(chan os.Signal)

	//	监听所有信号
	signal.Notify(p.stop, syscall.SIGINT, syscall.SIGTERM)

	//	阻塞数据直至信号传入
	<-p.stop

	//	关闭设备
	for _, boot := range p.boots {
		_ = boot.End()
	}
}

//	获取当前目录路径
//
//	Author(Wind)
func (p *Boost) GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Split(file)

	return path
}
