/**
 * Go 函数可以是一个闭包。闭包是一个函数值, 它引用了函数体之外的变量。
 * 这个函数可以对这个引用的变量进行访问和赋值;换句话说这个函数被 "绑定"
 * 在这个变量上。
 *
 * 例如, 函数 adder 返回一个闭包。每个返回的闭包都被绑定到其各自的 sum 变量上
 */
package main

import "fmt"


func fibonacci() func() int {

	back1, back2 := 0, 1
    return func() int {
        // 重新赋值
        back1, back2 = back2, (back1 + back2)
        return back1
    }
	
}

func main() {
	fib := fibonacci()
	for i := 1; i <= 10; i++ {
		fmt.Println(fib())
	}
}