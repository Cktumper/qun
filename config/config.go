package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

//	定义操作系统内核标识
const (
	CGOOSLinux   = "linux"
	CGOOSWindows = "windows"
	CGOOSMacOSX  = "darwin"
)

//	默认目录定义
const (
	DefaultUnixConfigPath    = "/etc"
	DefaultWindowsConfigPath = ""
)

type Config interface {
	Load() error
	Read(key string) string
}

var globalConfig Config

//	读取配置并转换为字符串
func ReadString(key string) string {
	return globalConfig.Read(key)
}

//	读取整型的配置文件
func ReadInt(key string) int {
	if value, err := strconv.Atoi(globalConfig.Read(key)); err == nil {
		return value
	}

	return 0
}

//	获取当前执行目录
func CurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return getDefaultETCDirectory()
	}

	return strings.Replace(dir, "\\", "/", -1)
}

//	获取默认的配置文件目录
func getDefaultETCDirectory() string {
	switch runtime.GOOS {
	case CGOOSWindows:
		return DefaultWindowsConfigPath
	}

	return DefaultUnixConfigPath
}
