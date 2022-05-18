package main

import (
	"fmt"
	"time"
)

// 程序创建之后会创建一个主 goroutine 去执行
func main() {
	for i := 0; i < 100000; i++ {
		go func() {
			fmt.Println("hello", i)
		}()
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main() 结束了 由 main 函数启动的 goroutine 也都结束了
}
