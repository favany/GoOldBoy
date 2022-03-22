package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // calc("1", 0, 3)
	a = 0
	defer calc("2", a, calc("20", a, b)) // calc("2", 0, 2)
	b = 1
}

// "10" 1 2 3
// "20" 0 2 2
// "2" 0 2 2
// "1" 0 3 3

//1. a := 1
//2. b := 2
//3. defer calc("1", 1, calc("10", 1, 2)) = defer calc("1", 1, 3)
//3.1  => calc("10", 1, 2) = "10" 1 2 3
//4. defer calc("2", 0, calc("20", 0, 2)) = defer calc("2", 0, 2)
//4.1  => calc("20", 0, 2) = "20" 0 2 2
// defer 后进先出，先执行后面的语句
//5. calc("2", 0, 2) = "2" 0 2 2
//6. calc("1", 1, 3) = "1" 1 3 4
