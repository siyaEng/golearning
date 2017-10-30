// 结构体
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) 
	// 调用结构体Vertex{1, 2}
}