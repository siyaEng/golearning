// 结构体字段可以通过结构体指针来访问。
// 通过指针间接的访问是透明的。
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // 指向 v
	p.X = 1e9 // 设置 X 的值
	fmt.Println(v)
}