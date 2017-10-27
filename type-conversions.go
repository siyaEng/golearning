//类型转换 T(v) 将 v 转换成类型 T
//
//var i int = 42
//var f float64 = float64(i)
//var u uint = uint(f)
//或者
//i := 42
//f := float(o)
//u := uint(f)
//与 c 不同 ，在类型之间项目赋值需要显示转换
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4 //移除var  int 或 float64 都不会有问题
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}