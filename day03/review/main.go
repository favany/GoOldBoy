package main

import "fmt"

// 复习
func main() {

	var name string
	name = "理想"
	fmt.Println(name)

	// 声明了一个变量，它是 [30]int 类型
	var ages [30]int
	ages = [30]int{1, 2, 3, 5, 6}
	fmt.Println(ages)

	// 自动匹配元素个数
	var ages2 = [...]int{1, 2, 3, 4}
	fmt.Println(ages2)

	// 指定元素索引
	var age3 = [...]int{1: 100, 10: 200}
	fmt.Println(age3)

	// 二维数组
	var a1 [3][2]int // [[1 2] [3 4] [5 6]]
	a1 = [3][2]int{
		{3, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

	// 数组是值类型
	x := [3]int{1, 2, 3}
	y := x     // 把 x 的值拷贝了一份给 y
	y[1] = 200 // 修改的是副本 y ，并不影响 x

	fmt.Println(x, y)

	// 切片（ Slice ）：
	var s1 []int
	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	var s2 []int
	fmt.Println(s2, s2 == nil) // 没有分配内存，== nil

	// make 必须要初始化 分配内存 不然没有内存
	s3 := make([]bool, 2, 4)
	fmt.Println(s3) // [false false]
	s4 := make([]int, 0, 4)
	fmt.Println(s4 == nil) // 分配内存， != nil

	c1 := []int{1, 2, 3}
	c2 := c1
	var c3 = make([]int, 3, 3)
	copy(c3, c1)

	fmt.Println(c2)
	c2[1] = 200
	fmt.Println(c2)
	fmt.Println(c1)

	// 指针
	// Go 里面的指针只能读，不能修改。不能修改指针变量指向的地址。
	addr := "沙河"
	addrP := &addr // 内存地址
	fmt.Println(addrP)
	fmt.Println("%T\n", addrP)
	addrV := *addrP // 根据内存地址找值
	fmt.Println(addrV)

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["lixiang"] = 10
	fmt.Println(m1)
	fmt.Println(m1["ji"]) // 如果key不存在，返回的是value对应类型的零值

	// 如果返回值是布尔值，我们通常用 ok 去接收它
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有这个人")
	} else {
		fmt.Println(score)
	}

	delete(m1, "li")           // 删除的key不存在 什么都不干
	delete(m1, "lixiang")      // 删除
	fmt.Println(m1, m1 == nil) // 已经开辟了内存，因此不为 nil
	
}
