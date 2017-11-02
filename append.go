// append 是 go 的内建函数，向 slice 末尾添加元素的操作
// 例子 func  append(s []T, vs ...T) []T 
// append 的第一个参数 s 是一个元素类型为 T 的 slice，其余类型为 T 的值会附加到
// slice 的末尾
// append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice 
// 如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。
// 返回的 slice 会指向这个新分配的数组
package main

import "fmt"

func main() {
	var a[]int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0) // 容量指的是什么 为什么会变化
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
} 	