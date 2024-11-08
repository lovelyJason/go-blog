package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type TomlConfig struct {
	Viewer Viewer // 虽然是大写，但是同toml里的能匹配上
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *TomlConfig

func init() {
	//	程序启动的就会执行init
	Cfg = new(TomlConfig)
	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	currentDir, err := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	_, err1 := toml.DecodeFile("config/config.toml", &Cfg)
	if err1 != nil {
		panic(err)
	}
}
