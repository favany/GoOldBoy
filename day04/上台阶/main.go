package main

import "fmt"

// 上台阶的面试题
// n 个台阶，一次可以走一步，也可以走两步，有多少种走法？

func step(n uint64) uint64 {
	// 如果只有一个台阶，就一种走法
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-1) + step(n-2)
}

func main() {
	ret := step(3)
	fmt.Println(ret)
}
