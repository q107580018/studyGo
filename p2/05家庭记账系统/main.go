package main

import (
	"fmt"
)

type MyAcount struct {
}

func Login() int {
	fmt.Println(`----------家庭收支记账软件----------
	
	  1 收支明细
	  2 登记收入
	  3 登记支出
	  4 退出

	`)
	var Myselect int
	fmt.Print("请选择(1-4):")
	fmt.Scan(&Myselect)
	return Myselect
}
func main() {
L:
	for {
		switch Login() {
		case 1:
		case 2:
		case 3:
		case 4:
			fmt.Println("您已退出家庭记账软件")
			break L
		default:
		}
	}

}
