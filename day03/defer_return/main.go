package main

import "fmt"

// Go 语言中函数的 return 不是原子操作，在底层是分为两步来执行
// 第一步，返回值赋值
// 第二步，函数中如果存在 defer ,那么执行 defer
// 第三步，真正的RET返回

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是 x 不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 先把返回值赋值为5，再执行defer，x 赋值为 6，因此最终返回值是 6

}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 第二步，修改的是 x，但是 y 没有被修改
	}()
	return x // 第一步，返回值是y = x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值是 5
}

func main() {
	valueF1 := f1()
	fmt.Println("f1:", valueF1)

	valueF2 := f2()
	fmt.Println("f2:", valueF2) // 6

	valueF3 := f3()
	fmt.Println("f3:", valueF3)

	valueF4 := f4()
	fmt.Println("f4:", valueF4)

}
