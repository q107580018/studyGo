package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(s)
	fmt.Println("------end------")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Printf("type:%v,value:%v,kind:%v\n", typ, val, kd)
	num := val.NumField() // 获取这个结构体有多少个字段
	fmt.Println("该结构体有", num, "个字段")
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d的值为%v\n", i, val.Field(i))
		// 获取struct的标签，注意需要使用reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d的tag为%v\n", i, tagVal)
		}
	}
}

// 发射遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
func main() {
	m := Monster{"qiu", 12, 100.0, "男"}
	TestStruct(m)
}
