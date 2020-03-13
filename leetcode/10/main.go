/*
面试题57 - II. 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：

1 <= target <= 10^5
*/
package main

func main() {

}
func sumF(a, b int) int {
	return ((a + b) * (b - a + 1) / 2)
}
func findContinuousSequence(target int) [][]int {
	var ans = [][]int{}
	start := 1
	end := 2
	for end-1 <= target/2 {
		sum := sumF(start, end)
		for sum > target {
			sum -= start
			start++
		}
		for sum < target {
			end++
			sum += end
		}
		if sum == target {
			a := []int{}
			for i := start; i <= end; i++ {
				a = append(a, i)
			}
			ans = append(ans, a)
			end++
			start = 1
			continue
		}

	}
	return ans
}
