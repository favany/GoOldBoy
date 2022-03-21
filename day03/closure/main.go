package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func sum(x, y int) {
	fmt.Println("this is sum()")
	fmt.Println(x + y)
}

// 定义函数包装器包装 funcWrapper
func funcWrapper(x func(int, int), m, n int) func() {
	temp := func() {
		x(m, n)
	}
	return temp
}

// 闭包
func main() {
	ret := funcWrapper(sum, 100, 200)
	ret()
}
