package bootstrap

import (
	"flag"
	"os"
)

type (
	//	私有接口，用于记录命令列表
	Cmd interface {
		//	注册命令
		Register(fg *flag.FlagSet)

		//	执行主体
		Handle() error

		//	命令名称
		Name() string
	}

	//	命令对象
	Command struct {
		fg     *flag.FlagSet
		cmd    []Cmd
		actual map[string]Cmd
	}
)

func NewCommand() *Command {
	p := &Command{
		fg:     flag.NewFlagSet("Run", flag.ContinueOnError),
		actual: make(map[string]Cmd),
	}

	//	返回自己
	return p
}

//	获取需要执行的命令列表
//
//	Author(Wind)
func (p *Command) GetShouldExecCommands() []Cmd {
	return p.cmd
}

//	注册命令
//
//	Author(Wind)
func (p *Command) Register(c Cmd) *Command {
	//	加入待执行的命令索引中
	p.actual[c.Name()] = c

	//	注册命令
	c.Register(p.fg)

	//	返回自己
	return p
}

//	判定是否是命令模式
//
//	Author(Wind)
func (p *Command) CommandMode() bool {
	return len(p.cmd) > 0
}

//	解析命令
//
//	Author(Wind)
func (p *Command) Parse() {
	if len(os.Args) > 1 && (os.Args[1] == "Run" || os.Args[1] == "run") {
		//	解析子命令
		_ = p.fg.Parse(os.Args[2:])

		//	获取执行参数
		p.fg.Visit(func(f *flag.Flag) {
			if cmd, ok := p.actual[f.Name]; ok {
				p.cmd = append(p.cmd, cmd)
			}
		})
	}
}
