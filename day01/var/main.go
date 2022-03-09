package main

import "fmt"

// Go 语言中推荐使用驼峰式命名

//var student_name string
//var studentName string
//var StudentName string

// 声明变量
var name string

var (
	age  int
	isOk bool
)

func main() {
	name = "理想"
	age = 16
	isOk = true

	// Go 语言中变量声明必须使用，不使用就编译不过去
	fmt.Print(isOk) // 在终端中
	fmt.Println()
	fmt.Printf("name:%s\n", name) // %s是占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)              // 打印完指定内容之后，会在后面加一个换行符

	// 声明变量，同时赋值
	var s1 string = "bingo"
	fmt.Println(s1)

	// 类型推导（根据值，判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明
	s3 := "哈哈哈"
	fmt.Println(s3)

	// 匿名变量
	// 注意事项：
	// 1. 函数外的每个语句都必须以关键字开始（var、const、func）
	// 2. :=不能使用在函数外
	// 3. 下划线_多用于占位，表示忽略值。（后面学了函数再说）
	// 4. 同一个作用域中，不能重复声明同名的变量
}
