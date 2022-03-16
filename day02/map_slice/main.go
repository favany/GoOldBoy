package main

import "fmt"

// map 和 slice组合

func main() {
	// 元素类型为 map 类型的切片
	var s1 = make([]map[int]string, 10, 10)

	//s1[0][100] = "A" // panic: assignment to entry in nil map 没有对内部的map做初始化

	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"
	fmt.Println(s1) // [map[10:沙河] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// 元素类型为 切片类型的 map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
