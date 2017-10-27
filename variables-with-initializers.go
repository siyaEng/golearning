package main

import "fmt"

var i, j int = 1, 2 // 变量类型从初始值中获得

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}