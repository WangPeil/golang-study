package gin_study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetParams(r *gin.Engine) {
	r.GET("/get_params", func(c *gin.Context) {
		// 第一种获取Get请求的参数
		first := c.Query("name")
		fmt.Printf("The first method get params: %s\n", first)
		// 第二种获取Get请求的参数
		second := c.DefaultQuery("name", "name")
		fmt.Printf("The second method get params: %s\n", second)
		// 第三种获取Get请求的参数
		third, _ := c.GetQuery("name")
		fmt.Printf("The third method get params: %s\n", third)
		c.JSON(http.StatusOK, gin.H{
			"get param first method":  first,
			"get param second method": second,
			"get param third method":  third,
		})
	})
}

func PostForm(r *gin.Engine) {
	r.POST("/post_forms", func(c *gin.Context) {
		// 第一种获取Post请求的参数
		first := c.PostForm("name")
		fmt.Printf("The first method post params: %s\n", first)
		// 第二种获取Post请求的参数
		second := c.DefaultPostForm("name", "name")
		fmt.Printf("The second method post params: %s\n", second)
		// 第三种获取Post请求的参数
		third, _ := c.GetPostForm("name")
		fmt.Printf("The third method post params: %s\n", third)
		c.JSON(http.StatusOK, gin.H{
			"post form first method":  first,
			"post form second method": second,
			"post form third method":  third,
		})
	})
}
