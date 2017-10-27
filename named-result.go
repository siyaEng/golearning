package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x
	return
	// 没有参数 return 语句返回各个变量的当前值。
	// 这种做法叫做 "裸" 返回
	// 这种返回语句用在短函数中。在长函数中会影响代码可读性
}

func main() {
	fmt.Println(split(17)) //返回结果是 int
}