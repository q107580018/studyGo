package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// BubbleSort 冒泡法
func BubbleSort(lists []int) {
	for i := 0; i < len(lists)-1; i++ {
		for j := 0; j < len(lists)-i-1; j++ {
			if lists[j] > lists[j+1] {
				lists[j], lists[j+1] = lists[j+1], lists[j]
			}

		}
	}
}

// selectionSort 选择排序
func selectionSort(l []int) {
	var minIndex int
	for i := 0; i < len(l)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(l); j++ {
			if l[j] < l[minIndex] {
				minIndex = j
			}
		}
		l[i], l[minIndex] = l[minIndex], l[i]
	}
}

// InsertionSort 插入排序
func InsertionSort(lists []int) {
	for i := 1; i < len(lists); i++ {
		tmp := lists[i]
		j := i - 1
		for ; j >= 0 && tmp < lists[j]; j-- {
			lists[j+1] = lists[j]
		}
		lists[j+1] = tmp
	}
}

// 快速排序
func quicksort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := 0
	r := len(nums) - 1
	base := nums[l]
	for {
		for {
			if l == r {
				nums[r] = base
				break
			}
			if nums[r] < base {
				nums[l] = nums[r]
				break
			}
			r--
		}
		for {
			if l == r {
				nums[l] = base
				break
			}
			if nums[l] > base {
				nums[r] = nums[l]
				break
			}
			l++
		}
		if l == r {
			break
		}
	}
	quicksort(nums[:l])
	quicksort(nums[l+1:])
}
func main() {
	rand.Seed(time.Now().UnixNano())
	l1 := func(i int) (m []int) {
		for k := 0; k < i; k++ {
			m = append(m, rand.Intn(100000))
		}
		return m
	}(100000)
	l2 := make([]int, len(l1))
	l3 := make([]int, len(l1))
	l4 := make([]int, len(l1))
	l5 := make([]int, len(l1))
	copy(l2, l1)
	copy(l3, l1)
	copy(l4, l1)
	copy(l5, l1)
	start := time.Now()
	BubbleSort(l1)
	// fmt.Println(l1)
	fmt.Println("冒泡法合计用时：", time.Since(start))
	start = time.Now()
	selectionSort(l2)
	// fmt.Println(l2)
	fmt.Println("选择法合计用时：", time.Since(start))
	start = time.Now()
	InsertionSort(l3)
	//fmt.Println(l3)
	fmt.Println("插入法合计用时：", time.Since(start))
	start = time.Now()
	quicksort(l4)
	//fmt.Println(l4)
	fmt.Println("快速排序法合计用时：", time.Since(start))
	start = time.Now()
	sort.Ints(l5)
	//fmt.Println(l5)
	fmt.Println("官方自带法合计用时：", time.Since(start))
}
