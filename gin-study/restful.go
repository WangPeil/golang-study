package gin_study

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestfulStart(r *gin.Engine) {
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodGet,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodPost,
		})
	})

	r.PUT("/put", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodPut,
		})
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodDelete,
		})
	})

	r.Any("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": c.Request.Method,
		})
	})
}
func PathVariable(r *gin.Engine) {
	r.GET("/path/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.Param("name"),
		})
	})

	r.GET("/path/:name/*action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":   c.Param("name"),
			"action": c.Param("action"),
		})
	})

}
