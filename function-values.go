/**
 * 函数也是值。他们可以像其他值一样传递, 比如,
 * 函数可以作为函数的参数或者返回值
 */
package main

import (
 	"fmt"
 	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(math.Pow)
	fmt.Println(compute(math.Pow))
}