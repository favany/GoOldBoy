package main

import "fmt"

// defer
// Go 中的 defer 语句会将其后面跟随的语句进行延迟处理。在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 定义的逆序进行执行。
// 也就是说，先被 defer 的语句最后执行，最后 defer 的语句最先执行。
// defer 多用于函数结束之前释放资源。

func deferFunc() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("呵呵呵")
	defer fmt.Println("哈哈哈")
	fmt.Println("end")
}

func main() {
	deferFunc()
}
