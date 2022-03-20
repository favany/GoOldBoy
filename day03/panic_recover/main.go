package main

import "fmt"

// panic 和 recover

func funcA() {
	fmt.Println("A")
}

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover() // recover 必须搭配 defer 使用
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}()
	panic("出现了严重的错误！！！") // 程序崩溃，退出
	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
