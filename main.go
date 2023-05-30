package main

import (
	"fmt"
	"canvine/setting"
	"canvine/router"
)

func main() {
	//1、加载配置
	if err := setting.Init(); err != nil {
		fmt.Println("init setting failed, ", err)
		return
	}

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Config.Port)); err != nil {
		fmt.Println("run server failed, err:", err)
		return
	}
}