package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//// 定义一个函数对 f2 进行包装
func lixiang(x func(int, int), m, n int) func() {
	temp := func() {
		x(m, n)
	}
	return temp
}

// 闭包
func main() {
	ret := lixiang(f2, 100, 200)
	ret()
}
