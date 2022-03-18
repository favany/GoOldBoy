package main

import "fmt"

// 变量的作用域

var x = 100 // 定义一个全局变量

// 定义一个函数
// 函数中查找变量的顺序
// 1. 现在函数内部查找
// 2. 找不到就往函数的外面查找，一直找到全局
// 3. 函数内部定义的变量只能再函数内部使用
// 4. if for switch 语句中定义的变量，也是只在 for 语句块中生效

func f1() {
	fmt.Println(x)
	//fmt.Println(y) // undefined: y
}

// 1.全局作用域
func main() {
	// 2. 函数块作用域

	f1()
	//fmt.Println(name) // 函数内部定义的变量只能再函数内部使用

	if i := 10; i < 18 { // 3. 语句块作用域
		fmt.Println("乖乖上学")
	}
	//fmt.Println(i) // undefined: i
}
