package main

import "fmt"

func printPercentage(num int) {
	fmt.Printf("%d%%\n", num)
}

func main() {
	fmt.Print("Meggy")
	fmt.Print("Go")
	fmt.Print("")
	fmt.Println("----------------------------------------------")
	fmt.Println("Meggy")
	fmt.Println("Go")
	// fmt.Printf("格式化字符串", 值 )
	// %T : 查看类型
	// %d : 十进制数
	// %b : 八进制数
	// %x : 十六进制数
	// %c : 字符
	// %s : 字符串
	// %q : 双引号括起来的字符串字面值
	// %p : 指针
	// %v : 值
	// %f : 浮点数
	// %% : 转义，表示正常的百分号

	var m1 = make(map[string]int, 1)
	m1["理想"] = 100
	fmt.Printf("%v", m1)  // map[理想:100]
	fmt.Printf("%#v", m1) // map[string]int{"理想":100}
	fmt.Printf("%q", "123")

	printPercentage(90)

	// 宽度标识符 % 宽度 . 精度 f ，宽度和精度可选填
	fmt.Println("\n宽度表示符")
	n := 12.3456
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%9.f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)

	// 获取输入 fmt.Scan()
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：", s)

	//var (
	//	name  string
	//	age   int
	//	class string
	//)
	//
	//fmt.Scanf("%s %d %s\n", &name, &age, &class)
	//fmt.Println(name, age, class)

}
