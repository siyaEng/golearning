package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) 
		// defer 是按照栈顺序的后进先出的原则;
		// 队列是先进先出的原则
	}

	fmt.Println("done")
}