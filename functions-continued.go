package main

import "fmt"

func add(x, y int) int {// x,y int 是 x int, y int的简写
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}