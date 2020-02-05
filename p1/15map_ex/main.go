package main

import (
	"fmt"
	"sort"
)

// 按顺序遍历map

func main() {
	myMap := make(map[string]string, 100)
	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("stu%02d", i)
		v := fmt.Sprintf("%c", i+13312)
		myMap[k] = v
	}
	keys := make([]string, 0, 100)
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(key, myMap[key])
	}

}
