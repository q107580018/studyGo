package main

/*
动态规划，求最优解
*/

import "fmt"

var prev = map[int]int{
	1: 0,
	2: 0,
	3: 0,
	4: 1,
	5: 0,
	6: 2,
	7: 3,
	8: 5,
}
var v = map[int]int{
	1: 5,
	2: 1,
	3: 8,
	4: 4,
	5: 6,
	6: 3,
	7: 2,
	8: 4,
}

var opt = make(map[int]int)

func main() {

	for i := 0; i < 9; i++ {
		opt[i] = max(i)
	}
	for i := 1; i < 9; i++ {
		v := opt[i]
		fmt.Printf("opt[%v]=%v\n", i, v)
	}

}
func max(i int) int {
	if (v[i] + opt[prev[i]]) > opt[i-1] {
		return v[i] + opt[prev[i]]
	} else {
		return opt[i-1]
	}
}
