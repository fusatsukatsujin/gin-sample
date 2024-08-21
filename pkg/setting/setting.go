package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	AppSetting.JwtSecret = cfg.Section("app").Key("JwtSecret").String()
}
