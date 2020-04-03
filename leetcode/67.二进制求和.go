/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	i := len(a)-1
	j := len(b)-1
	flag := 0
	res :=""
	for i>=0 || j>=0{
		intA , intB :=0,0
		if i>=0{
			if  a[i] == '1'{
				intA = 1
			}
		}
		if j>=0{
			if  b[j] == '1'{
				intB = 1
			}
		}
		count := intA+intB+flag
		flag = 0
		if count >= 2{
			count -=2
			flag = 1
		}
		switch count{
			case 0:
				res = "0"+res
			case 1:
				res = "1"+res
			}
		i--
		j--
	}
	if flag ==1 {
		res = "1"+res
	}	
	return res
	
}

// @lc code=end
