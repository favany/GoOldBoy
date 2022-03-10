package main

import "fmt"

// 流程控制之跳出for循环

func main() {
	// 当 i = 5 时跳出循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
	fmt.Println("over")

	// 当 i = 5 时跳过此次for循环（不执行 for 循环内部的打印语句）
	for i := 10; i < 10; i++ {
		if i == 5 {
			continue // 继续下一次循环
		}
	}
	fmt.Println("over")
}
