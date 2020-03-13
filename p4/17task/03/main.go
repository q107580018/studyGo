package main

import "fmt"

// 启动一个协程，将1-2000的数放入到一个channel中，比如numChan
// 启动8个协程，从numChan取出数(比如n),并计算1+...+n的值，并
// 存放到resChan
// 最后8个协程协同完成工作后，再遍历resChan，显示结果如[res[1]=1...]
// 注意，考虑resChan chan int是否合适?
var numChan = make(chan int)
var resChan = make(chan [2]int, 2000)
var isOkChan = make(chan bool)

func main() {
	resMap := make(map[int]int)
	go MakeChan()
	for i := 0; i < 8; i++ {
		go Sum()
	}
	for i := 0; i < 8; i++ {
		<-isOkChan
	}
	close(resChan)
	for res := range resChan {
		resMap[res[0]] = res[1]
	}
	for i := 1; i <= 2000; i++ {
		fmt.Printf("res[%v]=%v\n", i, resMap[i])
	}
}
func MakeChan() {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func Sum() {
	for num := range numChan {
		res := 0
		for i := 1; i <= num; i++ {
			res += i
		}
		resChan <- [2]int{num, res}
	}
	isOkChan <- true
}
