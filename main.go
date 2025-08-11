package main

import (
	"blog-admin/config"
	"blog-admin/core"
	"blog-admin/global"
	"blog-admin/router"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.DbInit()          // 初始化数据库
	core.LoggerInit()        // 初始化日志
	r := router.RouterInit() // 初始化路由
	global.Log.Warnf("当前运行环境为：%s", config.Config.App.Env)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.App.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Log.Errorf("ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP) // 使用无缓冲通道接收信号可能会导致信号丢失，此处虽无缓冲通道可工作，但为避免潜在问题，后续使用已创建的无缓冲通道即可
	<-quit
	global.Log.Warnf("正在关闭服务器 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 等待当前请求结束，最多等 5 秒
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Infoln("服务器强制退出:", err)
	}

	global.Log.Infof("服务器已优雅退出")
}
