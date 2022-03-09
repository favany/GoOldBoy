package main

import "fmt"

// 常量
// 定义了常量之后，不能修改
// 在程序运行期间，不会改变

// 批量声明常量时，如果某一行声明后没有赋值。默认就和上一行一致
const (
	statusOk = 200
	notFound = 404
)

//
const (
	n1 = 100
	n2
	n3
)
const pi = 3.1415926

func main() {
	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)

}
