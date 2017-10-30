// 类型 [n]T 是一个有 n 个类型为 T 的值的数组
// 表达式 var a [10]int
// 定义变量 a 是一个有十个整数的数组
// 数组不能改变大小
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) //打印数组元素
	fmt.Println(a) //打印数组
}