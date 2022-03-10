package main

import "fmt"

// if 条件判断

func main() {
	//age := 19
	//if age > 18 {
	//	fmt.Print("澳门首家线上赌场开业了！")
	//} else {
	//	fmt.Print("改写暑假作业了！")
	//}

	// 多个判断条件
	// age 变量此时只在 if 条件判断语句中生效
	if age := 19; age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

	//fmt.Println(age)
}
