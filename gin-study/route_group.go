package gin_study

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteGroup(r *gin.Engine) {
	// 路由组
	// 路由组中也可以添加新的路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "v1.add",
			})
		})
		security := v1.Group("/security")
		{
			security.GET("/add", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "v1.security.add",
				})
			})
		}
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "v2.add",
			})
		})
	}
}
