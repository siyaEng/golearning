// map 的文法跟结构体文法类似，不过必须有键名
package main

import "fmt"

type Test struct {
	bob, pop int
}

var m = map[string]Test{
	"first": {
		1, 2,
	},
	"second": {
		3, 4,
	},
}

func main() {
	for k, _ := range m {
		fmt.Println(m[k])
	}
}