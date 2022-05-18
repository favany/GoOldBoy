package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano()) // 保证每次执行的时候都不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) // 0 <= x < 10
		fmt.Println(-r1, -r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// 如何知道这十个 goroutine 都结束了
	wg.Wait()
}
