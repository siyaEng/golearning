package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum// 循环赋值给自己
		fmt.Println(sum)
	}
	//fmt.Println(sum)
}