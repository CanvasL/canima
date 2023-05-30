package main

import (
	"canvine/dao/mysql"
	"canvine/router"
	"canvine/settings"
	"canvine/utils/snowflake"
	"fmt"
)

func main() {
	// 1. init setting
	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed, err:", err)
		return
	}

	// 2. init mysql
	if err := mysql.Init(settings.Config.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}
	defer mysql.Close()

	// 3. init utils snowflake
	if err := snowflake.Init(settings.Config.StartTime, settings.Config.MachineID); err != nil {
		fmt.Println("init snowflake failed, err:", err)
	}

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Config.Port)); err != nil {
		fmt.Println("run server failed, err:", err)
		return
	}
}
