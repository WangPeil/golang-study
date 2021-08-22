package gin_study

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// auth 可以返回闭包
// 闭包中可以封装其他的信息 控制闭包中的环境
func auth(needAuth bool) gin.HandlerFunc {
	if needAuth {
		return func(c *gin.Context) {
			c.Set("auth", true)
		}
	}
	return func(c *gin.Context) {
		c.Set("auth", false)
	}
}

func Middleware(r *gin.Engine) {
	// 第一种添加中间件的方式
	r.Use(auth(true))
	r.GET("first_add_middleware", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"auth": c.GetBool("auth"),
		})
	})
	// 第二种添加中间件的方式
	// 第一种和第二种不能同时使用
	r.GET("second_add_middleware", auth(true), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"auth": c.GetBool("auth"),
		})
	})
}
