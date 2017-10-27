// go 的基本类型有Basic types

// bool
// string
// int int8 int16 int32 int64 有符号整型
// uint uint8 uint16 uint32 uint64 uintptr 无符号整型
// int, uint, uintptr 类型在 32 位的, 而在 64 位上试 64 位
// uintptr 无符号整型,用于存放一个指针
// byte // uint8 的别名
// rune // int32 的别名 代表一个Unicode码
// float32 float64
// complex64 32位实数和虚数 complex128 64位实数和虚数
package main

import (
	"fmt"
	"math/cmplx" // complex 数学函数包
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1 //位运算
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n" //和这个 const 一起用
	fmt.Printf(f, ToBe, ToBe) // printf打印变量类型和值
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}