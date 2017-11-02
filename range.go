// for 循环的 range 格式可以对 slice 和 map 进行迭代循环
// 当 for 使用循环遍历第一个 slice 时，每次 range 都会返回两个值。
// 第一个是当前下表序号，另一个是对当前元素的拷贝
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow{
		fmt.Printf("2**%d = %d\n", i, v)
	}
}