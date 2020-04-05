/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.64%)
 * Likes:    341
 * Dislikes: 0
 * Total Accepted:    111.8K
 * Total Submissions: 297.2K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
func mySqrt(x int) int {
return	dich(0,x,x)
}
func dich(head,tail,x int)int{
	if tail - head <=1{
		if tail * tail <= x{
			return tail
		}else{
			return head
		}
	}
	mid := (tail + head )/2
	if mid * mid <=x && (mid+1)*(mid+1) > x {
		return mid
	}else if mid * mid > x {
		return dich(head,mid,x)
	}else{
		return dich(mid,tail,x)
	}
}
// @lc code=end

