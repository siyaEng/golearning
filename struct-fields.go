// 访问结构体
// 结构体使用 . 来访问
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(Vertex{})
}