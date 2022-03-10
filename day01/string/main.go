package main

import (
	"fmt"
	"strings"
)

// 字符串：Go中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（ int、bool、float32、float64 等）一样。Go 里的字符串的内部实现 UTF-8 编码。字符串的值为 双引号"" 中的内容，可以在 Go 语言的源码中直接添加非 ASCII 码字符
// Go 语言中字符串是用双引号包裹的！！！
// Go 语言中单引号包裹的是字符！！！

// 字符串转义符

func main() {
	// \本来是具有特殊含义的，用\\转义，表示单纯的反斜线\
	path := "D:\\Go\\src\\day01"
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 多行的字符串用反引号
	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
		`
	fmt.Println(s2)

	// 字符串相关操作
	//fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	world := "dsb"

	ss := name + world
	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	// 分割
	ret := strings.Split(path, "\\")
	fmt.Printf("ret:%s\n", ret)

	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))

	// 前缀和后缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	fmt.Println(strings.HasSuffix(ss, "理想"))

	// 判断子串出现的位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))

}
