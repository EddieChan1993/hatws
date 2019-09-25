//+build !prod

package s

import (
	"log"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"github.com/go-ini/ini"
)

const RunMode = gin.DebugMode //调试模式
func load() {
	Cfg, err = ini.Load("conf/app.ini", "conf/app.dev.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	loadServer()
}
func loadServer() {
	Service = new(Server)
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	Service.HTTPAdd = sec.Key("HTTP_ADDR").MustString("")
	Service.HTTPPort = sec.Key("HTTP_PORT").MustString("8000")
	Service.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	Service.WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	fmt.Println("server is running in 【开发模式】")
}
