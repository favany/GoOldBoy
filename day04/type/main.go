package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型
type myInt int     // 自定义类型：在编译后也保留
type yourInt = int // 类型别名：只在代码编写过程中有效，编译后不保留

func main() {
	var n myInt
	var m yourInt
	n = 100
	m = 100
	fmt.Println(n)        // 100
	fmt.Printf("%T\n", n) // main.myInt
	fmt.Println(m)        // 100
	fmt.Printf("%T\n", m) // int

	var c rune
	c = '中'
	fmt.Println(c)        // 20013
	fmt.Printf("%T\n", c) // int32
}
