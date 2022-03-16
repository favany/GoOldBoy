package main

import "fmt"

// map
func main() {
	// 声明 map 类型
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化（没有在内存中开辟空间）
	// map 的初始化
	m1 = make(map[string]int, 10) // 要估算好该map的容量，避免在程序运行期间 动态扩容

	m1["理想"] = 18
	m1["jiwuming"] = 35

	fmt.Println(m1)
	fmt.Println(m1["理想"])

	fmt.Println(m1["娜扎"]) // 如果不存在这个 key ，拿到对应值类型的零值

	// 约定俗成 拿ok
	value, ok := m1["娜扎"]
	fmt.Println(value, ok)
	if !ok {
		fmt.Println("查无此key", value, ok)
	} else {
		fmt.Println(value, ok)
	}

	// map 的遍历
	fmt.Println("map 的遍历：")
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	fmt.Println("只遍历 key：")
	// 只遍历 key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历 value
	fmt.Println("只遍历 value：")
	for _, value := range m1 {
		fmt.Println(value)
	}

	// 使用delete() 函数删除键值对 delete(map, key)
	fmt.Println("使用delete() 函数删除键值对")
	delete(m1, "jiwuming")
	fmt.Println(m1)

	// 删除一个不存在的key，不报错
	delete(m1, "bingo")
	fmt.Println(m1)

}
