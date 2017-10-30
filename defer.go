package main

import "fmt"

func main() {
	defer fmt.Println("world")// defer 延迟函数执行到上层返回

	fmt.Println("hello")
}