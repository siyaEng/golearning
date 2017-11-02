// 可以通过 _ 渠道 range 的序号和值
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(" pow = ", pow)
	for i := range pow{
		pow[i] = 1 << uint(i) // 迭代赋值
	}

	fmt.Println(" pow = ", pow)

	for _, value := range pow{
		fmt.Printf("%d\n", value)
	}

}