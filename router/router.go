package router

import (
	"github.com/gin-gonic/gin"
	"go-api/handler/sd"
	"go-api/handler/user"
	"go-api/router/middleware"
	"net/http"
)

// 配置文件解析、路由注册及加载及响应、添加中间件
// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 设置 HTTP Header
	// 数通过 g.Use() 来为每一个请求设置 Header，

	// 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，
	// 这时候为了不影响下一次请求的调用，需要通过 gin.Recovery() 来恢复 API 服务 器
	g.Use(gin.Recovery())
	g.Use(gin.Logger())       // 输出到控制台的 请求log
	g.Use(middleware.NoCache) // 强制浏览器不使用缓存
	g.Use(middleware.Options) // 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Secure)  // 一些安全设置
	g.Use(mw...)
	// 路由url 错误情况 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//用于身份验证功能的API
	g.POST("/login", user.Login)

	//新增一个创建用户的 API
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware()) //api鉴权
	{
		//u.POST("", user.Create)
		//u.POST("/:username", user.Create)

		u.POST("", user.Create1)      //创建用户
		u.DELETE("/:id", user.Delete) //删除用户
		u.PUT("/:id", user.Update)    //更新用户
		u.GET("", user.List)          //用户列表
		u.GET("/:username", user.Get) //获取指定用户的详情

	}

	// 定义了一个叫 sd 的分组，在该分组下注册了
	// /health 、 /disk 、 /cpu 、 /ram HTTP 路径，
	// 分别路由到 sd.HealthCheck 、 sd.DiskCheck 、 sd.CPUCheck 、 sd.RAMCheck 函
	//sd 分组主要用来检查 API Server 的状态：健康状况、服务器硬盘、CPU 和内存使用量
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck) //健康检查
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("ram", sd.RAMCheck)
	}
	return g
}
