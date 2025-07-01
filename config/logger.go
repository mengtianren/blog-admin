package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"blog-admin/global"

	"github.com/sirupsen/logrus"
)

// InitLogger 初始化日志配置
func InitLogger() {
	// 初始化日志实例
	global.Log = logrus.New()

	// 获取日志配置
	loggerConfig := Config.Logger

	// 设置日志级别
	level, err := logrus.ParseLevel(loggerConfig.Level)
	if err != nil {
		level = logrus.DebugLevel // 默认调试级别
	}
	global.Log.SetLevel(level)

	// 设置日志格式
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableQuote:    true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			if loggerConfig.ShowLine {
				filename := filepath.Base(f.File)
				return "", fmt.Sprintf(" [%s:%d]", filename, f.Line)
			}
			return "", ""
		},
	}
	global.Log.SetFormatter(formatter)

	// 创建日志目录
	if err1 := os.MkdirAll(loggerConfig.Director, 0755); err1 != nil {
		global.Log.Fatalf("创建日志目录失败: %v", err1)
	}

	// 设置日志输出
	logFile := filepath.Join(loggerConfig.Director, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		global.Log.Fatalf("打开日志文件失败: %v", err)
	}

	// 输出到控制台和文件
	if loggerConfig.LogInConsole {
		global.Log.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		global.Log.SetOutput(file)
	}

	// 设置显示行号
	if loggerConfig.ShowLine {
		global.Log.SetReportCaller(true)
	}

	global.Log.Info("日志初始化成功")
}
