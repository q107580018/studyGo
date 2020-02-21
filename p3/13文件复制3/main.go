package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CopyFile ...
func CopyFile(srcFileName, dstFileName string) {
	// 判断是否同名
	if srcFileName == dstFileName {
		fmt.Println("目标文件名不能和源文件名相同")
		return
	}
	readByte, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		fmt.Println("ioutil.ReadFile err:", err)
		return
	}
	err = ioutil.WriteFile(dstFileName, readByte, 0666)
	if err != nil {
		fmt.Println("ioutil.WriteFile err:", err)
		return
	}

}

func main() {
	list := os.Args
	if len(list) < 3 {
		fmt.Println("格式：xxx srcFile dstFile")
		return
	}
	sF := list[1]
	dF := list[2]
	CopyFile(sF, dF)
}
