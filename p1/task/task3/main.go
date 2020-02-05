// 判断 101-200 之间有多少个素数，并输出所有素数。
package main

import "fmt"

func main() {
	var prime = make([]int, 0, 100)
	for i := 101; i < 201; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			prime = append(prime, i)
		}

	}
	fmt.Printf("101-200之间一共有%d个素数\n", len(prime))
	fmt.Println(prime)
}
