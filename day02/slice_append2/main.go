package main

import "fmt"

// 关于 append 删除切片中的某个元素

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17} // 数组
	s1 := a1[:]                                   // 切片

	// 删除索引为 1 的那个 3
	s1 = append(s1[:1], s1[2:]...)

	fmt.Println(s1) // [1 5 7 9 11 13 15 17]
	fmt.Println(a1) // [1 5 7 9 11 13 15 17 17]

}
