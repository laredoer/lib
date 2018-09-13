package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("获取url中的域名")

	a := "http://www.waylau.com/golang-strings-split-get-url/"

	a1 := strings.Split(a, "//")[1]
	a2 := strings.Split(a1, "/")[0]
	str := "http://" + strings.Split(strings.Split(a, "//")[1], "/")[0]
	fmt.Println(a1)  //输出为：www.waylau.com/golang-strings-split-get-url/
	fmt.Println(a2)  //输出为：www.waylau.com
	fmt.Println(str) //输出为：http://www.waylau.com
}
