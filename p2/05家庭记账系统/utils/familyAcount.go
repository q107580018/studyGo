package utils

import (
	"bufio"
	"fmt"
	"os"
)

type MyAcount struct {
	Details []string
	Balance float64
	IsQuit  bool
}

func NewMyAcount() *MyAcount {
	return &MyAcount{Balance: 10000}
}

func Login() int {
	fmt.Println(`
----------家庭收支记账软件----------
	
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

func (m *MyAcount) IAndEDetals() {
	fmt.Println("\n----------当前收支明细记录----------")
	fmt.Println("收支\t账户金额\t收支金额\t说    明\t")
	if len(m.Details) == 0 {
		fmt.Println("当前没有收支明细...来一笔吧！")

	} else {
		for _, v := range m.Details {
			fmt.Print(v)
		}
	}

}

func (m *MyAcount) IncomeRegistrate() {
	fmt.Print("本次收入金额:")
	var i float64
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("输入错误。。。")
		return
	}
	fmt.Print("本次收入说明:")
	r := bufio.NewReader(os.Stdin)
	note, _ := r.ReadString('\n')
	m.Balance += i
	m.Details = append(m.Details, fmt.Sprintf(
		"收入\t%-7.1f\t\t%-9.1f\t%s", m.Balance, i, note))
	fmt.Println("\n--------------登记完成--------------")

}
func (m *MyAcount) ExpenseRegistrate() {
	fmt.Print("本次支出金额:")
	var i float64
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("输入错误。。。")
		return
	}
	if i > m.Balance {
		fmt.Println("余额不足...")
		return
	}
	fmt.Print("本次支出说明:")
	r := bufio.NewReader(os.Stdin)
	note, _ := r.ReadString('\n')
	m.Balance -= i
	m.Details = append(m.Details, fmt.Sprintf(
		"支出\t%-7.1f\t\t%-9.1f\t%s", m.Balance, i, note))
	fmt.Println("\n--------------登记完成--------------")

}

func (m *MyAcount) Exit() {
	for {
		flag := false
		fmt.Print("你确定要退出吗？y/n:")
		var k string
		fmt.Scan(&k)
		switch k {
		case "y", "Y":
			fmt.Println("您已退出家庭记账软件")
			m.IsQuit = true
			return
		case "n", "N":
			flag = true
		default:
			fmt.Println("您的输入有误，请重新输入 y/n")
		}
		if flag {
			break
		}

	}
}

func (m *MyAcount) MainMenu() {
	for {
		switch Login() {
		case 1:
			m.IAndEDetals()
		case 2:
			m.IncomeRegistrate()
		case 3:
			m.ExpenseRegistrate()
		case 4:
			m.Exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if m.IsQuit {
			break
		}
	}

}
