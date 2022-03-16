package main

import "fmt"

// 指针 pointer
// 指针
// Go 语言不存在指针操作，只需要记住两个符号：
// 1. & 取地址
// 2. * 根据地址取值

func main() {
	n := 18
	fmt.Println(&n) // 0x1400018c008
	p := &n
	fmt.Printf("%T\n", p) // *int: int 类型的指针
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) // int

	//var a *int
	//*a = 100
	//fmt.Println(*a) // panic: runtime error: invalid memory address or nil pointer dereference

	var a1 *int     // nil pointer
	fmt.Println(a1) // nil

	// new() 函数申请一个地址

	var a2 = new(int)
	fmt.Println(a2)  // 0x1400018c020
	fmt.Println(*a2) // 0
	*a2 = 100
	fmt.Println(*a2) // 100

	// make 和 new 的区别
	// 1. make 和 new 都是用来申请内存的
	// 2. new 很少用，一般用来给基本数据类型申请内存，string\int ...，返回的是对应类型的指针（*string、 *int）
	// 3. make 是用来给 切片slice 、map 、 channel 申请内存的，make 函数返回的是这三个类型本身

}

// 总结：
// 取地址操作符 & 和取值操作符 * 是一对互补操作符，& 取出地址 ， * 根据地址取出地址指向的值
// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
// - 对变量进行取地址操作（&），可以获得这个变量的指针变量
// - 指针变量的值是指针地址
// - 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
