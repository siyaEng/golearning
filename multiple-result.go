package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x // 结果返回两个字符串
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}