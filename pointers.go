// Go 具有指针。指针保存了变量的内存地址
// 类型 *T 是执行类型 T 的值的指针。其零值是 nil。
// var p *int 定义一个指针 p
// & 符号会生成一个指向其作用对象的指针
// i := 42
// p = &i
// * 符号表示指针指向的底层的值。
// fmt.Println(*p) // 通过指针 p 读取 i
// *p = 21 // 通过指针 p 设置 i
// 这也就是通常所说的 "间接饮用" 或 "或非直接饮用"。
// 与 C 不同, Go 没有指针运算
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // p 指向 i
	fmt.Println(*p) // 读取 i 通过 指针 p
	fmt.Println(i)

	p = &j // p 指向 j
	*p = *p / 37 // 设置了 j 的值 通过指针 p
	fmt.Println(j) 
}