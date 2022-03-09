package main

import "fmt"

// iota 是go语言的常量计数器，只能在常量的表达式中使用。
// iota 在const关键字出现时，将被重置为0。
// const 中每新增一行常量声明，将使iota计数一次（iota可理解为const语句块中的行索引）。iota类似枚举。
// 使用iota能简化定义，在定义枚举时很有用。
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2 = iota // 1
	_         // 2
	b3        // 3
)

const (
	c1 = iota // 0
	c2 = 100  // 1
	c3 = iota // 2
	c4        // 3
)

// 多个变量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1: 1 d2: 2 iota = 0
	d3, d4 = iota + 1, iota + 2 // d3: 2 d4: 3 iota = 1
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)

}
