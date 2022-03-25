package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func getGreet(f func(string)) {
	f("bingo")
}

func sum() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

// 闭包
func closure(f func(string), str string) func() {
	return func() {
		f(str)
	}
}

func main() {
	getGreet(greet)
	getSum := sum()
	sumResult := getSum(3, 4)
	fmt.Println(sumResult)

	// 执行闭包
	runClosure := closure(greet, "bingo")
	runClosure()
}
