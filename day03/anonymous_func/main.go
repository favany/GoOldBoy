package main

// 匿名函数
import "fmt"

var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)

	// 函数内部没办法声明带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果只调用一次函数，还可以简写成立即执行的函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("Hello world!")
	}(100, 200)

}
