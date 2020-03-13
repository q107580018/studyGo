package main

import (
	"fmt"
	"reflect"
)

// 用于演示反射
func TestReflect(b interface{}) {
	// 通过反射，获取传入变量的type，kind，值
	// 1.先获取type
	rType := reflect.TypeOf(b)
	fmt.Println("Type:", rType)
	// 2.获取到value
	rVal := reflect.ValueOf(b)
	fmt.Println("Value:", rVal)
	// num:= 2 + rVal 这样写是不能用的，rVal的实际类型是reflect.Value
	num := 2 + rVal.Int()
	fmt.Println("num:", num)
	// 下面将 rVal转为interface{}
	iV := rVal.Interface()
	// 将 interface 通过类型断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}
func main() {
	var num int = 10
	TestReflect(num)
}
