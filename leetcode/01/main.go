package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

func longestCommonPrefix(strs []string) string {
	buf := make([]byte, 0)
	minIndex := 0
	if len(strs) == 0 {
		return ""
	}
	for i, v := range strs {
		if v == "" {
			return ""
		}
		if len(v) < len(strs[minIndex]) {
			minIndex = i
		}
	}
	for i := 0; i < len(strs[minIndex]); i++ {
		fmt.Println("i=", i)
		for i2, v := range strs {
			if i2 == minIndex {
				continue
			}
			if strs[minIndex][i] != v[i] {
				return string(buf)
			}
		}
		fmt.Println("添加", string(strs[minIndex][i]))
		buf = append(buf, strs[0][i])
		fmt.Println("buf=", buf)
	}
	return (string(buf))
}
func main() {
	a := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println("a=", a)
}
