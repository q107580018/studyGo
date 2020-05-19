/*
414. 第三大的数
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

示例 1:

输入: [3, 2, 1]

输出: 1

解释: 第三大的数是 1.
示例 2:

输入: [1, 2]

输出: 2

解释: 第三大的数不存在, 所以返回最大的数 2 .
示例 3:

输入: [2, 2, 3, 1]

输出: 1

解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。
*/
package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(thirdMax(nums))
}
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	Maxs := make([]int, 0)
	for _, v := range nums {
		switch {
		case len(Maxs) == 0:
			Maxs = append(Maxs, v)
		case len(Maxs) == 1:
			if v > Maxs[0] {
				Maxs = []int{v, Maxs[0]}
			} else if v < Maxs[0] {
				Maxs = append(Maxs, v)
			}
		case len(Maxs) == 2:
			if v > Maxs[0] {
				Maxs = []int{v, Maxs[0], Maxs[1]}
			} else if v < Maxs[1] {
				Maxs = append(Maxs, v)
			} else if v < Maxs[0] && v > Maxs[1] {
				Maxs = []int{Maxs[0], v, Maxs[1]}
			}
		case len(Maxs) == 3:
			if v > Maxs[0] {
				Maxs = []int{v, Maxs[0], Maxs[1]}
			} else if v < Maxs[0] && v > Maxs[1] {
				Maxs = []int{Maxs[0], v, Maxs[1]}
			} else if v < Maxs[1] && v > Maxs[2] {
				Maxs = append(Maxs[:2], v)
			}
		}
	}
	if len(Maxs) < 3 {
		return Maxs[0]
	} else {
		return Maxs[2]
	}
}
