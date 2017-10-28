package main

import "fmt"

func main() {
	sum := 0 //初始值,循环体用大括号包起来
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}