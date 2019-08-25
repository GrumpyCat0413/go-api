package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/router"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("hello world!")
	//gin.SetMode(gin.ReleaseMode)

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
		log.Print("The router has been deployed successfully.") //路由器已成功部署
	}()

	log.Printf("start to listening the incoming requests on http address %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
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
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			// 服务器可用
			return nil
		}

		log.Print("Waiting for the router,retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router.")
}
