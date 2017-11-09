/**
 * 接受者为指针的方法
 * 方法可以与命名类型或命名类型的指针关联
 */
package main

import (
 	"fmt"
 	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Befote scaling: %+v, Abs: %v\n", v, v.Abs()) // %+v可以原样输出结构体
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}