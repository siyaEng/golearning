// map 的文法跟结构体文法类似，不过必须有键名
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {

	for k, v := range m {
		fmt.Printf("%v's postion is %v\n", k, v)
	}
}