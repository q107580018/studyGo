/*
面试题59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/
package main

func main() {

}

type MaxQueue struct {
	s   []int
	max []int
}

func Constructor() MaxQueue {

	return MaxQueue{make([]int, 0), make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.s) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.s = append(this.s, value)
	if len(this.max) == 0 {
		this.max = append(this.max, value)
		return
	}
	if value >= this.max[0] {
		this.max = []int{value}
		return
	}
	for i := len(this.max) - 1; this.max[i] < value; i-- {
		this.max = this.max[:i]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.s) == 0 {
		return -1
	}
	PopInt := this.s[0]
	this.s = this.s[1:]
	if PopInt == this.max[0] {
		this.max = this.max[1:]
	}
	return PopInt
}
