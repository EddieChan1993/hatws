package main

import (
	"hatgo/pkg/link"
	"fmt"
	"hatgo/pkg/conf"
	"log"
	"hatgo/pkg/logs"
	"hatgo/app/router"
)

const keyVer = "[version]"
var _version_ = "none setting"

func main() {
	defer func() {
		link.Db.Close()
		link.Rd.Close()
		logs.LogsReq.Close()
		logs.LogsSql.Close()
		logs.LogsWs.Close()
	}()

	router := router.InitRouter()
	log.Printf("%s %s",keyVer,_version_)
	err := router.Run(fmt.Sprintf("%s%s", conf.Serverer.HTTPAdd, conf.Serverer.HTTPPort))
	if err != nil {
		log.Fatalf("[server stop]%v", err)
	}
}
