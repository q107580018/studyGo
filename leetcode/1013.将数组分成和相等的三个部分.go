/*
1013. 将数组分成和相等的三个部分
给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。



示例 1：

输出：[0,2,1,-6,6,-7,9,1,2,0,1]
输出：true
解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
示例 2：

输入：[0,2,1,-6,6,7,9,-1,2,0,1]
输出：false
示例 3：

输入：[3,3,6,5,-2,2,5,1,-9,4]
输出：true
解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4


提示：

3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test")
	r := bufio.NewReader(f)
	r.ReadSlice(',')
	a := []int{1, -1, 1, -1}
	fmt.Println(canThreePartsEqualSum(a))
}
func canThreePartsEqualSum(A []int) bool {
	if len(A) <= 2 {
		return false
	}
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	sum1 := 0
	for i := 0; i < len(A)-2; i++ {
		sum1 += A[i]
		if sum1 == sum/3 {
			sum2 := 0
			for k := len(A) - 1; k > i+1; k-- {
				sum2 += A[k]
				if sum1 == sum2 {
					return true
				}
			}
			return false
		}

	}
	return false
}
