package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	res := 0
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
