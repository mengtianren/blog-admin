package main

import (
	"blog-admin/config"
	"blog-admin/core"
	"blog-admin/global"
	"blog-admin/router"
	"fmt"
)

func main() {
	config.DbInit()          // 初始化数据库
	core.LoggerInit()        // 初始化日志
	r := router.RouterInit() // 初始化路由
	global.Log.Warnf("当前运行环境为：%s", config.Config.App.Env)
	r.Run(fmt.Sprintf(":%d", config.Config.App.Port))
}
