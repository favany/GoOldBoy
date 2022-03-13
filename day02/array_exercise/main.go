package main

import "fmt"

// array 数组练习题
// [1, 3, 5, 7, 8]
// 1. 求数组所有元素的和
// 2. 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8 的两个元素的下标分别为(0, 3)和(1, 2)

func main() {
	// 1
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)

	// 2
	// 定义 2个for 循环，外层从第一个开始遍历，内层从外层后一个元素开始遍历

	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}

	}
}
