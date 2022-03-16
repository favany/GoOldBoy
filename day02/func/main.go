package main

import "fmt"

// 函数

// 函数存在的意义
// 函数是一段代码的封装。
// 把一段逻辑抽象出来，封装到一个函数中，给它起个名字。每次用到它的时候，直接调用即可。
// 使用函数能使代码结构更清晰，更简洁。

// 函数的定义
func sum(x int, y int) (sumResult int) {
	return x + y
}

// 没有返回值
func sum1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数、没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数，但有返回值
func f3() int {
	return 3
}

// 返回值可以命名，也可以不命名
// 命名的返回值就相当于在函数中声明了一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值，return 后可以省略
}

// 多个返回值
func f5() (int, int) {
	return 1, 2
}

// 参数的类型简写：当参数中，连续多个参数的类型一致时，我们可以将前面那个参数的类型省略
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y 的类型是切片 []int
}

// Go 语言中函数没有默认参数这个概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	sum1(1, 2)

	f2()

	f3()

	m, n := f5()
	println(m, n)

	f6(1, 2, 3, "5", "string", true, false)
}
