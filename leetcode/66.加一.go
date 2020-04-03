/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	add := false
	for i:=len(digits)-1;i>=0;i--{
		if add {
			digits[i]++
		}
		if digits[i]>9 && i==0{
			digits[i]= digits[i]%10
			digits = append([]int{1},digits...)
			break
		}else if digits[i]>9{
			digits[i]= digits[i]%10
			add = true
		}else{
			break
		}

	}
	return digits
}
// @lc code=end

