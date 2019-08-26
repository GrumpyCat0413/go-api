package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-api/config"
	"go-api/router"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "go apiserver config file path.")
)

// cfg变量从命令行flag 传入，可以传值 比如：./go-api -c config.yaml
// 如果为空会默认读取conf/config.yaml

func main() {
	fmt.Println("hello world!")

	pflag.Parse() // 命令行解析 cfg=配置文件的路径

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	var middlewares []gin.HandlerFunc

	// 配置文件解析
	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response , or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.") //路由器已成功部署
	}()

	log.Infof("start to listening the incoming requests on http address: %s", viper.GetString("addr")) //从配置文件读取配置地址
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// API 服务器健康状态自检
/*
有时候 API 进程起来不代表 API 服务器正
API 进程存在，但 是服务器却不能对外提供服务
在启动 HTTP 端口前 go 一个 pingServer 协程， 启动 HTTP 端口后，
该协程不断地 ping /sd/health 路径，如果失败次数超过一定次数，则终 止 HTTP 服务器进程。
通过自检可以最大程度地保证启动后的 API 服务器处于健康状态。
*/
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		//resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			// 服务器可用
			return nil
		}

		log.Info("Waiting for the router,retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router.")
}
