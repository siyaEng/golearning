// 对 slice 切片
// slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
// 表达式 s[lo:hi]
// 表示从 lo 到 hi-1 的 slice 元素，含前端，不包含都断。
// 因此 s[lo:lo] 是空的，而 s[lo：lo+1] 有一个元素
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下表代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])

	// 取所有
	fmt.Println("s[:] ==", s)
	fmt.Println("s[:] ==", s[0:len(s)])
}