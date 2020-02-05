package main

import "fmt"

type dog struct {
	name string
}

// 构造函数
func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

// 方法是作用于特殊函数
func (d *dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}

func (d *dog) setName(name string) {
	d.name = name
}

func main() {
	d1 := newDog("金毛")
	d1.wang()
	d1.setName("萨摩")
	fmt.Println(d1.name)
}
