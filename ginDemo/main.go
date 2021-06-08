package main

import (
	"fmt"
	"ginDemo/routers"
	"ginDemo/app/dataParseBind"
)

func main() {
	// 创建路由
	// 加载多个APP的路由配置
	routers.Include(dataParseBind.Routers)
	// 初始化路由
	r := routers.Init()
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
