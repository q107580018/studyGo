package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(FileName string) (NewFileName string) {
		if strings.HasSuffix(FileName, suffix) {
			NewFileName = FileName

		} else {
			NewFileName = FileName + suffix
		}
		return
	}
}

func main() {
	h := makeSuffix(".jpg")
	fmt.Println(h("aaa.txt"))
	fmt.Println(h("123"))
	fmt.Println(h("1.jpg"))
}
