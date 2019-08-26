package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
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

	// 初始化日志包
	c.initLog()

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
		log.Infof("Config file changed: %s", e.Name)
	})
}

// 日志初始化配置
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)

}

/*
log 配置文件参数：
writers: 输出位置  可选项 file和stdout 选择file将日志记录到logger_file指定的文件中，选择 stdout 会将日志输出到标准输出，当然也可以两者
logger_level: 日志级别 DEBUG、INFO、WARN、ERROR、FATAL
logger_file :  日志文件
log_format_text : 日志的输出格式，JSON 或者 plaintext， true 会输出成非 JSON 格式， false 会输出成 JSON 格式
rollingPolicy : rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如 果是 size 则根据大小进行转存
log_rotate_date : rotate 转存时间，配 合 rollingPolicy: daily 使用
log_rotate_size : rotate 转存大小，配合 rollingPolicy: size 使用
log_backup_count : 当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份， 这里指定了备份文件的大个数
*/
