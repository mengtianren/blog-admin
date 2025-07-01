package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Config *appConfig

type appConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port int    `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"app"`
	// 日志
	Logger struct {
		Level        string `mapstructure:"level"`
		Prefix       string `mapstructure:"prefix"`
		Director     string `mapstructure:"director"`
		ShowLine     bool   `mapstructure:"show_line"`
		LogInConsole bool   `mapstructure:"log_in_console"`
	} `mapstructure:"logger"`
	// 数据库
	Database struct {
		Host      string `mapstructure:"host"`
		Port      int    `mapstructure:"port"`
		User      string `mapstructure:"user"`
		Password  string `mapstructure:"password"`
		DBName    string `mapstructure:"dbname"`
		Charset   string `mapstructure:"charset"`
		ParseTime string `mapstructure:"parseTime"`
		Loc       string `mapstructure:"loc"`
	} `mapstructure:"database"`
}

func init() {
	// 1. 获取环境变量，默认 dev
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	configName := fmt.Sprintf("config.%s", strings.ToLower(env))
	fmt.Printf("configName: %v\n", configName)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&Config)
	fmt.Printf("✅ YAML 配置加载成功:\n%+v\n", Config.Logger)
}
