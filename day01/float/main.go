package main

import (
	"fmt"
)

// 浮点数
func main() {
	//math.MaxFloat32 // 32位浮点数最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认Go中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显式声明float32类型
	//f1 = f2                // float32 类型的值不能直接赋值给 float64 类型

}

// 布尔值：Go语言中以 bool 类型声明布尔型数据，布尔型数据只有 true 和 false 两个值。
//
