package main

import "fmt"

func main() {
	for count := 1; count < 10; count++ {
		for i := 1; i <= count; i++ {
			fmt.Print(count, " * ", i, " = ", count*i)
			fmt.Print("   ")
		}
		fmt.Println("")
	}
}
