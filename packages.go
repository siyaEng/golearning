package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())//rand.Seed种子生成其,生产随机数
	fmt.Println("My favorite number is", rand.Intn(10))// rand.Intn随机函数
}