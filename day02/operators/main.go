package main

import "fmt"

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)

	// 算术运算符

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// ++ 和 -- 是单独的语句，不是运算符
	a++
	b--
	fmt.Println(a)

	// 关系运算符: Go是强类型的语言，相同类型的变量才能比较。
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	// 逻辑运算符
	// && ：如果年龄大于18 并且小于60
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班")
	} else {
		fmt.Println("不用上班！")
	}

	// || ：如果年龄大于18 或者 小于60
	if age < 18 || age > 60 {
		fmt.Println("不用上班！")
	} else {
		fmt.Println("苦逼上班！")
	}

	// ! ：取反，原来为真，就为假
	isMarried := false
	println(isMarried, !isMarried)

	// 位运算符：针对二进制
	// 5的二进制表示: 101
	// 2的二进制表示:  10

	// &:按位与 - 参与运算的两数各对应的二进制位相与。（两位均为1才为1）
	fmt.Println(5 & 2) // 000 -> 0
	// |:按位或 - 参与运算的两数各对应的二进制位相或。（两位有1个1就为1）
	fmt.Println(5 | 2) // 111 -> 7
	// ^:按位异或 - 按位异或。（两位不一样则为1）
	fmt.Println(5 ^ 2) // 111 -> 7
	// <<: 将二进制位数左移制定位数
	fmt.Println(5 << 1)  // 1010 -> 10
	fmt.Println(1 << 10) // 10000000000 -> 1024
	// >>: 将二进制位数右移制定位数
	fmt.Println(5 >> 1) // 10 -> 2

	var m = int8(1)
	fmt.Println(m << 10) // 'm' (8 bits) too small for shift of 10

	// 192.168.1.1
	// 权限 文件操作会有位运算实际的应用
	// 0644

	// 赋值运算符： 用来给变量赋值

	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // x = x - 1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2

	x <<= 2 // x = x << 2
	x >>= 2 // x = x >> 2
	x &= 2  // x = x & 2
	x |= 3  // x = x | 3
	x ^= 4  // x = x ^ 4

}
