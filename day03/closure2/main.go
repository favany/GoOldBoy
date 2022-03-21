package main

// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的变量。

// 底层的原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找。

import "fmt"

func adder() func(int) int {
	x := 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	// 闭包等于下方函数+外部变量的引用x
	return func(y int) int {
		x += y
		return x
	}
}

func f3(x, y int) func() {
	return func() {
		fmt.Println(x + y)
	}
}

func main() {
	ret := adder()
	result := ret(200)
	fmt.Println(result)
	runF3 := f3(1, 2)
	runF3()
}
