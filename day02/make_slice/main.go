package main

import "fmt"

// make函数创建切片

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d \n", s1, len(s1), cap(s1))

	// 切片的本质
	// 切片就是一个框，框住了一块连续的内存。
	// 属于引用类型，真正的数据保存在底层数组中。
	// 切片之间不能直接比较。只能和 nil 比较。
	// 一个 nil 值切片的长度和容量都是0。
	// 但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例
	// var s1 []int // len(s1)=0; cap(s1)=0; s1==nil
	// s2 := []int{} // len(s2)=0; cap(s2)=0; s2!=nil
	// s3 := make([]int, 0)

	// 所以要判断一个切片是否为空，要用len(s) == 0 来判断，不应该使用 s == nil 来判断

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 // s3 和 s4 都指向同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	// 切片的遍历
	// 1. 索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// 2. for range 循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
