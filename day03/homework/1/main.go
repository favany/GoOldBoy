package main

import (
	"fmt"
	"unicode"
)

// 判断字符串中汉字的数量
func main() {
	// 难点是判断一个字符是汉字
	s1 := "Hello沙河"
	var count int
	// 1. 依次拿到字符串中的字符
	for _, c := range s1 {
		// 2. 判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			// 3. 把汉字出现的次数累加得到最终结果
			count++
		}
	}
	fmt.Println(count)

}
