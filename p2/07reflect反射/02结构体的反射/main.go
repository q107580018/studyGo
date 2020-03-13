package main

import (
	"fmt"
	"reflect"
)

// 结构体的反射

type Student struct {
	Name string
	age  int
}

func reflectTest(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType:", rTyp)
	rVal := reflect.ValueOf(b)
	fmt.Println("rValue:", rVal)
	IV := rVal.Interface()
	stu, ok := IV.(Student) // 可以通过switch类型断言，更灵活
	if ok {
		fmt.Println(stu.Name)
	}
}

func main() {
	stu := Student{"qiu", 18}
	reflectTest(stu)
}
