package main

import "fmt"

// 复习
func main() {

	var name string
	name = "理想"
	fmt.Println(name)

	// 声明了一个变量，它是 [30]int 类型
	var ages [30]int
	ages = [30]int{1, 2, 3, 5, 6}
	fmt.Println(ages)

	// 自动匹配元素个数
	var ages2 = [...]int{1, 2, 3, 4}
	fmt.Println(ages2)

	// 指定元素索引
	var age3 = [...]int{1: 100, 10: 200}
	fmt.Println(age3)

	// 二维数组
	var a1 [3][2]int // [[1 2] [3 4] [5 6]]
	a1 = [3][2]int{
		{3, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

	// 数组是值类型
	x := [3]int{1, 2, 3}
	y := x     // 把 x 的值拷贝了一份给 y
	y[1] = 200 // 修改的是副本 y ，并不影响 x

	fmt.Println(x, y)

}