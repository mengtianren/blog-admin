package core

import (
	"blog-admin/config"
	"blog-admin/global"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func LoggerInit() {
	logger := logrus.New()

	// 设置日志等级
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// 设置格式
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		DisableQuote:    true,
		PadLevelText:    true,
	})
	logger.SetReportCaller(config.Config.Logger.ShowLine)

	// 设置文件输出（自动切割，按天备份）
	logFile := &lumberjack.Logger{
		Filename:   config.Config.Logger.Director + "/" + config.Config.Logger.Prefix + ".log", // 文件名自动带日期
		MaxSize:    100,                                                                        // 每个日志文件最大 100MB
		MaxBackups: 7,                                                                          // 最多保留 7 个备份
		MaxAge:     30,                                                                         // 最多保存 30 天
		Compress:   true,                                                                       // 是否压缩
	}

	// 根据环境分输出
	if config.Config.App.Env == "dev" {
		logger.SetOutput(os.Stdout)
	} else if config.Config.Logger.LogInConsole {
		logger.SetOutput(io.MultiWriter(logFile, os.Stdout))
	} else {
		logger.SetOutput(logFile)
	}

	global.Log = logger
}
