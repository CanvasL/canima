package main

import (
	"fmt"
	"canvine/setting"
	"canvine/router"
	"canvine/dao/mysql"
)

func main() {
	// 1. init setting
	if err := setting.Init(); err != nil {
		fmt.Println("init setting failed, err:", err)
		return
	}

	// 2. init mysql
	if err := mysql.Init(setting.Config.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}
	defer mysql.Close()

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Config.Port)); err != nil {
		fmt.Println("run server failed, err:", err)
		return
	}
}