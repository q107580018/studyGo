package main

import (
	"errors"
	"fmt"
)

// MyDiv 除法运算
func MyDiv(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return result, err

}

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println("result = ", result)
	}
}
