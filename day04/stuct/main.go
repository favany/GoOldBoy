package main

import "fmt"

// 结构体

type person struct {
	name, gender string
	age          int
	hobby        []string
}

func main() {
	// 声明一个person类型的变量p
	var p person

	// 通过字段赋值
	p.name = "bingo"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"🏀", "⚽️", "双色球"}
	fmt.Println(p)

	// 访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Printf(p.name)
}
