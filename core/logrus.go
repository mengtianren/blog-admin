package core

import (
	"blog-admin/config"
	"blog-admin/global"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggerInit() {

	logger := logrus.New()
	// 设置日志等级
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	// 如果出错 默认使用info等级
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	// 设置输出格式
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false, // 确保不禁用颜色
		DisableQuote:    true,
		PadLevelText:    true,
	})
	// 是否显示调用行
	logger.SetReportCaller(config.Config.Logger.ShowLine)

	// 创建日志目录 如果不存在 则创建
	if _, err1 := os.Stat(config.Config.Logger.Director); os.IsNotExist(err1) {
		_ = os.MkdirAll(config.Config.Logger.Director, os.ModePerm)
	}

	// 日志文件名：按天分隔
	logName := fmt.Sprintf("%s-%s.log", config.Config.Logger.Prefix, time.Now().Format("2006-01-02"))
	filePath := path.Join(config.Config.Logger.Director, logName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件失败: %v", err))
	}

	//  如果是开发环境 不打印日志
	if config.Config.App.Env == "dev" {
		logger.SetOutput(io.MultiWriter(os.Stdout))
	} else {
		if config.Config.Logger.LogInConsole {
			logger.SetOutput(io.MultiWriter(file, os.Stdout))
		} else {
			logger.SetOutput(file)
		}
	}

	global.Log = logger
}
