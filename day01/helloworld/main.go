// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

package main

// 导入语句
import "fmt"

// 函数外只能放标识符（变量、常量、函数、类型）的声明

// 程序的入口函数
func main() {
	fmt.Println("Hello World!")
}
