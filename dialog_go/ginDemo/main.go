package main

import (
	"fmt"
	"ginDemo/app/dataParseBind"
	"ginDemo/app/middlewares"
	"ginDemo/app/shop"
	"ginDemo/routers"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	// 日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 创建路由
	// 加载多个APP的路由配置
	routers.Include(dataParseBind.Routers, shop.Routers)
	// 初始化路由并注册中间件
	r := routers.Init()
	// 注册中间件: 如果路由没有进行拆分的时候写在main.go, 拆分后写在routers.go中
	//r.Use(middlewares.GlobalMiddleWare())
	/**
	cookie 测试
	模拟实现权限验证中间件
	有2个路由，login和home
	login用于设置cookie
	home是访问查看信息的请求
	在请求home之前，先跑中间件代码，检验是否存在cookie
	访问home，会显示错误，因为权限校验未通过
	**/
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(http.StatusOK, "login success")
	})
	r.GET("/home", middlewares.AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "home",
		})
	})
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

	// 基础写法
	// 初始化路由
	//r := gin.Default()
	//// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello world")
	//})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	//r.Run(":8000")
}
