package dataParseBind

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*数据解析和绑定*/
// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

/*json 数据解析与绑定*/
func JSONHandler(c *gin.Context) {
	var json Login
	// 将request的body中的数据，自动按照接送格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 判断用户名密码
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user or password is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

/*表单 数据解析与绑定*/
func FormHandler(c *gin.Context) {
	var form Login
	// 将request的body中的数据，自动按照接送格式解析到结构体
	if err := c.Bind(&form); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 判断用户名密码
	if form.User != "root" || form.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user or password is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

/*表单 数据解析与绑定*/
func UriHandler(c *gin.Context) {
	var login Login
	// 将request的body中的数据，自动按照接送格式解析到结构体
	if err := c.ShouldBindUri(&login); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 判断用户名密码
	if login.User != "root" || login.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user or password is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
