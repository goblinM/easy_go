package shop

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	shoppingGroup := e.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
}
