package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

//	配置结构
type Config struct {
}

//	构建个新的配置文件管理器
//
//	Author(Wind)
func NewConfig() *Config {
	return &Config{}
}

//	加载配置文件
//
//	Author(Wind)
func (p *Config) Load(path string) {
	//	加载 ENV 文件
	if err := godotenv.Load(fmt.Sprintf("%s.env", path)); err != nil {
		log.Fatalln(err)
	}
}
