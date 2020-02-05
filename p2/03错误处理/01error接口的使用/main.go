package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建error接口方式一
	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println(err1)
	// 创建error接口方式二
	err2 := errors.New("this is normal err2")
	fmt.Println(err2)

}
