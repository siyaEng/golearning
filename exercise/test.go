package main

import (
	"fmt"
	"reflect"
)

const (
	b   = 'A' // rune 类型, int32 别名
	num = b * 20
)

func main() {
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b)) // 输出变量的类型
	fmt.Println(num)
}
