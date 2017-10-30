package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// time.Now()获取当前的日期
	// time.Now().Hour()获取当前的小时
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}