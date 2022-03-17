package main

import "strings"

// 2. how do you do 单词出现的次数

func main() {
	s2 := "how do you do"
	// 2.1 把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	// 2.2 遍历切片存储到一个 map
	m1 := make(map[string]int, 10)
	// 2.3 累加出现的次数
	for _, w := range s3 {
		m1[w]++
	}

	for key, value := range m1 {
		println(key, value)
	}
	
}

//// 1. 如果 map 中不存在 w 这个 value，这个 key 出现的次数 = 1
//if _, ok := m1[w]; !ok {
//	m1[w] = 1
//	// 2. 如果 map 中存在 w 这个 key ，那么出现的次数 + 1
//} else {
//	m1[w]++
//}
