package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
)

func main() {
	rename("/Volumes/home/Drive/学习教程/golang教程/beego", `_\d\d_`, "")
}

func rename(filePath, pattern, repName string) {
	filesInfo, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal("读取路径错误")
	}

	for _, fileInfo := range filesInfo {
		fileName := fileInfo.Name()
		fmt.Printf("fileName=%v\n", fileName)
		reg, _ := regexp.Compile(pattern)
		newName := reg.ReplaceAllString(fileName, repName)
		fmt.Printf("newName=%v\n", newName)
		err := os.Rename(path.Join(filePath, fileName), path.Join(filePath, newName))
		if err != nil {
			fmt.Println("重命名失败", err)
		} else {
			fmt.Println("重命名成功")
		}
	}

}
