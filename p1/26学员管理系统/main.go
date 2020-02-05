package main

import (
	"fmt"
	"os"
)

// 学生类型
type student struct {
	ID   int
	name string
}

// 系统类型
type sms struct {
	allStudent map[int]*student
}

// 构造新学生的函数
func newStudent(ID int, name string) *student {
	return &student{
		ID:   ID,
		name: name,
	}
}

func (s sms) showAllStudent() {
	fmt.Println("-------------------------")
	for _, value := range s.allStudent {
		fmt.Printf("学员ID:%d 学员名字:%s\n", value.ID, value.name)
	}
}

func (s sms) addStudent() {
	var ID int
	var name string
	fmt.Print("请输入学员ID：")
	fmt.Scanln(&ID)
	fmt.Print("请输入学员名字：")
	fmt.Scanln(&name)
	student := newStudent(ID, name)
	s.allStudent[ID] = student
	fmt.Println("添加完毕")

}

func (s sms) deleteStudent() {
	var deleteID int
	fmt.Print("请输入要删除学员的ID：")
	fmt.Scanln(&deleteID)
	delete(s.allStudent, deleteID)
}

func (s sms) renameStudent() {
	var (
		renameID int
		newName  string
	)
	fmt.Print("请输入需要修改的学员的ID：")
	fmt.Scanln(&renameID)
	stu, ok := s.allStudent[renameID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("您需要修改的学员ID：%d, 姓名:%s\n", stu.ID, stu.name)
	fmt.Print("请输入新名字:")
	fmt.Scanln(&newName)
	stu.name = newName
	fmt.Println("修改成功")
}

func main() {
	var choice int
	allStudent := make(map[int]*student)
	s1 := sms{allStudent}
	for {
		fmt.Println("欢迎光临学员管理系统")
		fmt.Println("************************\n1.查看所有学员\n2.添加学员\n3.删除学员\n4.修改学员\n5.退出系统")
		fmt.Print("请输入指令:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			s1.showAllStudent()
		case 2:
			s1.addStudent()
		case 3:
			s1.deleteStudent()
		case 4:
			s1.renameStudent()
		case 5:
			fmt.Println("已退出，欢迎下次光临")
			os.Exit(1)

		default:
			fmt.Println("输入错误，请重新输入～")
		}
	}
}
