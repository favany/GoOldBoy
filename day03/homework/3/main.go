package main

import (
	"fmt"
	"strings"
)

// 回文判断：字符串从左往右读，和从右往左读是一样的，那么就是回文。
// 上海自来水来自海上
// 山西煤运车运煤西山
// 黄山落叶松叶落山黄
// 解题思路：

func palindrome(str string) (isPalindrome bool) {
	isPalindrome = true
	arr := make([]string, 100, 100)
	arr = strings.Split(str, "")
	len := len(arr)

	fmt.Println(arr)

	for i := 0; i < len; i++ {
		if arr[i] != arr[len-i-1] {
			isPalindrome = false
		}
	}

	return
}

func main() {
	test1 := palindrome("上海自来水来自海上")
	test2 := palindrome("山西煤运车运煤西山")
	test3 := palindrome("黄山落叶松叶落山黄")

	fmt.Println(test1, test2, test3)
}
