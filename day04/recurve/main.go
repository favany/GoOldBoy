package main

import "fmt"

// 永远不要高估自己

// 递归：自己调用自己
// 适合处理问题相同/问题规模越来越小的场景

// 3! = 3*2*1 = 3*2!
// 4! = 4*3*2*1 = 4*3!
// 5! = 5*4*3*2*1 = 5*4!

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return f(n-1) * n
}

func main() {
	ret := f(5)
	fmt.Println(ret)
}
