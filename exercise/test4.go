// golang 的位运算
package main

import (
	"fmt"
)

func main() {
	/*LABLE1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABLE1
			}
		}
	}*/

	/*	for {
			for i := 0; i < 10; i++ {
				if i > 3 {
					goto LABLE1
				}
			}
		}
	LABLE1:*/

	/*for {
			for i := 0; i < 10; i++ {
				if i > 3 {
					continue LABLE1
				}
			}
		}
	LABLE1:*/
	fmt.Println("OK")
}
