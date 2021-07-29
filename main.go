package main

import "fmt"

func main() {
	//fmt.Println("Hello world")
	//mysql_study.F1()
	//mysql_study.Query()
	//mysql_study.QueryAll()
	g()
}
func g() {
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
