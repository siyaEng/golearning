package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I am learning go!"

	slice := strings.Fields(s)

	for k, v := range slice {
		fmt.Println(k)
		fmt.Println(v)
	}

}