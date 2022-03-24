package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	// 1. 依次拿到每个人的名字
	for _, user := range users {
		// 2。 拿到一个人名根据金币的规则去分金币
		for _, letterInt := range user {
			// https://www.cnblogs.com/aaronthon/p/10787176.html
			// 字节 转换为 字符串
			letter := string(letterInt)
			// 名字中每包含1个'e'或'E'分1枚金币
			if strings.ToLower(letter) == "e" {
				distribution[letter] += 1
				// 名字中每包含1个'i'或'I'分2枚金币
			} else if strings.ToLower(letter) == "i" {
				distribution[letter] += 2
				// 名字中每包含1个'o'或'O'分3枚金币
			} else if strings.ToLower(letter) == "o" {
				distribution[letter] += 3
				// 名字中每包含1个'u'或'U'分4枚金币
			} else if strings.ToLower(letter) == "u" {
				distribution[letter] += 4
			}
		}
	}

	for _, usedCoin := range distribution {
		coins -= usedCoin
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
