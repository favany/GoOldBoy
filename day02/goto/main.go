package main

import "fmt"

// goto

func main() {
	var flag = false
	// goto 常用的方式：跳出多层 for 循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				// 设置退出标签
				break // 跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break // 跳出外层的for循环
		}
	}
}

//// goto + label 实现跳出多层 for 循环
//func main() {
//	for i := 0; i < 10; i++ {
//		for j := 'A'; j < 'Z'; j++ {
//			if j == 'B' {
//				// 设置退出标签
//				goto breakTag
//			}
//			fmt.Printf("%v-%c", i, j)
//		}
//	}
//	return
//
//	// 标签
//breakTag:
//	fmt.Println("结束for循环")
//}
