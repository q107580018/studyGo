/*
322. 零钱兑换
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。
*/
package main

import "fmt"

func main() {
	a := coinChange([]int{1, 2, 5}, 11)
	fmt.Println(a)
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		// fmt.Println("i=", i)
		dp[i] = -1
		for _, c := range coins {
			// fmt.Println("c=", c)
			if i < c {
				// fmt.Printf("i=%v, c=%v, i<c, continue\n", i, c)
				continue
			} else if dp[i-c] == -1 {
				// fmt.Printf("i=%v,c=%v,dp[i-c](dp[%v]) == -1, continue\n", i, c, i-c)
				continue
			}
			count := dp[i-c] + 1
			// fmt.Printf("count = dp[i-c]+1 = %v\n", count)
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
				// fmt.Printf("dp[i]=dp[%v]=count=%d\n", i, count)
			}
		}
	}
	return dp[amount]
}
