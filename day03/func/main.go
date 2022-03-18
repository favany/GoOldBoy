package main

import "fmt"

// 函数：一段函数的封装 📦

func f1() {
	fmt.Println("Hello 沙河！")
}

func f2(name string) {
	fmt.Println("Hello", name)
}

// 带参数和返回值的函数
func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(y) // y 是一个 int 类型的 切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	// 由于已声明了返回值 sum， 那么在函数中可以直接使用返回值变量
	sum = x + y
	return // 由于已声明了返回值 sum， return 后面可以省略返回值的变量
}

// Go 函数支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("理想")
	f2("姬无命")
	valueF3 := f3(100, 200)
	fmt.Println(valueF3)

	f5("1", 2, 3, 4)
	fmt.Println(f6(1, 2))

	// 在一个命名的函数中，不能再声明一个命名的函数
	//func f8 () {
	//
	//}
}
