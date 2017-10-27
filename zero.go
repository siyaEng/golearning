package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//变量定义时没有明确的初始化会赋值为零值
	//int => 0, bool => false, string => ""
}