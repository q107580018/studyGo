package main

import "fmt"

// Student ...
type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 10
	i[1] = "hello"
	i[2] = Student{"小明", 18}
	for _, data := range i {
		if v, ok := data.(int); ok == true {
			fmt.Println("该变量为int,内容为:", v)
		} else if v, ok := data.(string); ok == true {
			fmt.Println("该变量为string,内容为:", v)
		} else if v, ok := data.(Student); ok == true {
			fmt.Println("该变量为Student,内容为:", v)
		}

	}
}
