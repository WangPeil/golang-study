package main

import (
	"github.com/gin-gonic/gin"
	"golang-study/gin-study"
)

func main() {
	//fmt.Println("Hello world")
	//mysql_study.F1()
	//mysql_study.Query()
	//mysql_study.QueryAll()
	gin.DisableConsoleColor()
	r := gin.Default()
	gin_study.GetParams(r)
	gin_study.PostForm(r)
	gin_study.RouteGroup(r)
	gin_study.Middleware(r)
	_ = r.Run()
}
