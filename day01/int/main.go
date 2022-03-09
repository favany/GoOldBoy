package main

import "fmt"

// 整型
// 整型分为以下两大类：按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
// 其中，uint8 就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的Long型

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 把十进制转换为二进制
	fmt.Printf("%o\n", i1) // 把十进制转换为八进制
	fmt.Printf("%x\n", i1) // 把十进制转换为十六进制

	// 八进制 以0开头
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制 以0x开头
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9) // 明确制定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}
