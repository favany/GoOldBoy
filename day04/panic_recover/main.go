package main

import "fmt"

func f1() {
	defer func() {
		err := recover() // 收集当前的错误
		fmt.Println("去爱...")
		fmt.Println(err)

	}()
	panic("程序错误")
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
