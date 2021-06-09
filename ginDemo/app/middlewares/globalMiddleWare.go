package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**全局中间件*/
func GlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func GlobalUseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 统计时间
		since := time.Since(start)
		fmt.Println("程序用时：", since)
	}
}

/**验证cookie*/
func AuthMiddleWare() gin.HandlerFunc  {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				// 验证无误往下执行
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "auth is invalid",
		})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}