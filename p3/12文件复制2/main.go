package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CopyFile ...
func CopyFile(srcFileName, dstFileName string) {
	// 判断是否同名
	if srcFileName == dstFileName {
		fmt.Println("目标文件名不能和源文件名相同")
		return
	}
	// 打开源文件
	f, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	f2, err := os.Create(dstFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 完成后自动关闭
	defer f.Close()
	defer f2.Close()

	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(f2)
	for {
		readStr, err := reader.ReadString('\n')
		writer.WriteString(readStr)
		if err == io.EOF {
			break
		}
	}
	writer.Flush()

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
