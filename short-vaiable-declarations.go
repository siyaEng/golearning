package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3// := 是 var 的简写不能用在函数外
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}