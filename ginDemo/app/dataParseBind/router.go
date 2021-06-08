package dataParseBind

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine)  {
	e.POST("loginJSON", JSONHandler)
	// 校验： curl http://localhost:8000/loginJSON -H 'content-type:application/json' -d {"user":"root", "password":"admin"} -X POST
	// curl http://localhost:8000/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\", \"password\":\"admin\"}" -X POST
	// curl http://localhost:8000/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\", \"password\":\"admin2\"}" -X POST
	e.POST("loginForm", FormHandler)
	// 校验： 需要构造一个前端页面去校验
	// curl http://localhost:8000/loginForm -H 'application/x-www-form-urlencoded' -d "{\"user\":\"root\", \"password\":\"admin\"}" -X POST
	e.GET("/:user/:password", UriHandler)
	// 校验： curl http://localhost:8000/root/admin

}

