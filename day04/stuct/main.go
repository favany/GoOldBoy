package main

import "fmt"

// 结构体 P51

type person struct {
	name, cnName string
	age          int
	hobby        []string
}

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	// 声明 结构体
	// 实例化 person 类型 p
	var p person

	// 用 new 关键字对结构体进行实例化，得到结构体的指针
	var p2 = new(person)
	p2.cnName = "六"
	// 分别获取类型 值 和 内存地址（指针）
	fmt.Printf("type: %T value: %v pointer: %v \n", p2, p2, &p2)

	// 声明并初始化结构体
	var p3 = person{
		name:   "floy",
		cnName: "bu",
		age:    20,
		hobby:  []string{"🍰", "🍪"},
	}
	fmt.Println(p3)

	// 使用列表的方式初始化，要求排列的顺序和结构体声明时的顺序一致
	var p4 = person{"mophia", "jun", 22, []string{"🎱", "⚽️", "🏎️"}}
	fmt.Println(p4)

	// 通过字段赋值
	p.name = "mophia"
	p.age = 22
	p.cnName = "jun"
	p.hobby = []string{"🎱", "⚽️", "🏎️"}

	// 结构体的类型和值
	fmt.Printf("type:%T value:%v \n", p, p)
	// main.person {mophia jun 22 [🎱 ⚽️ 🏎️]}

	// 访问变量p的字段
	fmt.Printf(p.name)
	// mophia

	// 匿名结构体 anonymous struct ：多用于一些临时场景
	var aStruct struct {
		name string
		age  int
	}

	aStruct.name = "mophia"
	aStruct.age = 22
	fmt.Printf("type:%T value:%v \n", aStruct, aStruct)
	// type:struct { name string; age int } value:{Meggy 23}

	// 结构体的修改方法：指针传递. 推荐使用结构体指针
	changeAge(&p)
	fmt.Println(p.age)

	// 结构体占用一块连续的内存空间
	m := x{int8(10), int8(10), int8(10)}
	fmt.Printf("%p %p %p", &(m.a), &(m.b), &(m.c))
	// 0x1400012c010 0x1400012c011 0x1400012c012

	// 补充材料：内存对齐 https://segmentfault.com/a/1190000017527311

}

func changeAge(p *person) {
	// 可通过结构体指针访问结构体中的成员
	p.age += 1
}
