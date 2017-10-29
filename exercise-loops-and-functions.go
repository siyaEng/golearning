package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {// for if 循环跳出

	var result float64
	for i := 1; i <= 10; i++ {
		z := float64(i);
		z = z - (z*z - x) / 2*z

		if (math.Sqrt(x) - z) < 1.0 {
			result = z
			tmp = i
			break
		}
	}
	return result, tmp
}

func main() {
	//Sqrt(2)
	fmt.Println(Sqrt(2))
}