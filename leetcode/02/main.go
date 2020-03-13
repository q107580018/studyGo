package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	c := make([]rune, 0)
	MyMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, v := range s {
		if _, ok := MyMap[v]; ok {
			if len(c) == 0 {
				return false
			}
			if MyMap[v] == c[len(c)-1] {
				c = c[:len(c)-1]
			} else {
				return false
			}
		} else {
			c = append(c, v)
		}
	}

	if len(c) == 0 {
		return true
	} else {
		return false
	}

}
func main() {
	fmt.Println(isValid(""))
}
