package main

import (
	//"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	stringsplits := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range stringsplits {
		m[v] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}