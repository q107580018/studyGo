/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (66.03%)
 * Likes:    277
 * Dislikes: 0
 * Total Accepted:    70.1K
 * Total Submissions: 106K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 * 
 * 
 * 
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 * 
 * 示例:
 * 
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	if numRows == 2 {
		return [][]int{[]int{1}, []int{1, 1}}
	}
	res := make([][]int, 2, numRows)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		tmp[len(tmp)-1] = 1
		for j := 1; j < len(tmp)-1; j++ {
			tmp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, tmp)
	}
	return res
}

// @lc code=end

