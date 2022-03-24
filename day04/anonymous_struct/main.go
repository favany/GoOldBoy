package main

import "fmt"

// 匿名结构体：多用于一些临时场景

func main() {

	var anonymousStruct struct {
		name string
		age  int
	}

	anonymousStruct.name = "Meggy"
	anonymousStruct.age = 23

	fmt.Printf("type:%T value:%v", anonymousStruct, anonymousStruct)
	// type:struct { name string; age int } value:{Meggy 23}

}
