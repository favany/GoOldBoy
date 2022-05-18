package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序创建之后会创建一个主 goroutine 去执行
func main() {
	for i := 0; i < 100000; i++ {
		go hello(i) // 开启一个单独的 goroutine 去执行 hello 函数（任务）
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main() 结束了 由 main 函数启动的 goroutine 也都结束了

}
