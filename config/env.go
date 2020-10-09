package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

//	构建 Env 配置结构
type Env struct {
	Path string
}

//	初始化 Env 配置
func NewEnv(name string) Config {
	globalConfig = &Env{Path: fmt.Sprintf("%s/%s", CurrentDirectory(), name)}

	return globalConfig
}

//	加载配置数据
func (p *Env) Load() error {
	return godotenv.Load(p.Path)
}

//	读取配置
func (p *Env) Read(key string) string {
	return os.Getenv(key)
}
