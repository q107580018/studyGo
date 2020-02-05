package main

import "fmt"

// Student ...
type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 4)
	i[0] = 10
	i[1] = "hello"
	i[2] = Student{"小明", 18}
	i[3] = true

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Println("i[", index, "]该变量为int,内容为:", value)
		case string:
			fmt.Println("i[", index, "]该变量为string,内容为:", value)
		case Student:
			fmt.Println("i[", index, "]该变量为Student,内容为:", value)
		default:
			fmt.Println("i[", index, "]该变量为未知类型")
		}
	}
}
