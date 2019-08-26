package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//配置文件
type Config struct {
	Name string
}

func Init(cfg string) error {
	fmt.Println("init cfg：", cfg)
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件 并解析
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 配置文件的初始化函数 并解析配置文件
func (c *Config) initConfig() error {
	if c.Name != "" { //设置指定的配置文件
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf")   // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config") // 默认指定配置文件路径，以及设置配置文件名字
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML

	// 环境变量设置读取
	viper.AutomaticEnv()                      // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER")           // 读取环境变量的前缀为 APISERVER
	replacer := strings.NewReplacer(".", "_") // . 替换成 _
	viper.SetEnvKeyReplacer(replacer)
	//例如： export APISERVER_ADDR=:7777
	//export APISERVER_URL=http://127.0.0.1:7777

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序 不需要重启
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
