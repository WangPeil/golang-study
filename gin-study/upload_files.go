package gin_study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func SingleFileUpload(r *gin.Engine) {
	// 上传文件规定最大的缓冲区大小
	r.MaxMultipartMemory = 8 << 20
	r.POST("/single_upload", func(c *gin.Context) {
		// 单个文件
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		// 上传文件到指定的位置
		cwd, _ := os.Getwd()
		_ = c.SaveUploadedFile(file, path.Join(cwd, "file.txt"))
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}

// MultiFilesUpload 多文件上传
// curl -X POST http://localhost:8080/multi_upload -F "upload[]=@/Users/appleboy/test1.zip" -F "upload[]=@/Users/appleboy/test2.zip" -H "Content-Type: multipart/form-data"
func MultiFilesUpload(r *gin.Engine) {
	// 上传文件规定的最大缓冲区的大小
	r.MaxMultipartMemory = 8 << 20
	r.POST("multi_upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		cwd, _ := os.Getwd()
		for _, file := range files {
			fmt.Println(file.Filename)
			// 上传文件到指定路径
			_ = c.SaveUploadedFile(file, path.Join(cwd, file.Filename))
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
