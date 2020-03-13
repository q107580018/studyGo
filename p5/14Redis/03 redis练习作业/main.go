// 1.Monster信息[name age skill]
// 2.使用golang操作redis，存放3个monster到redis中。[比如使用hash数据类型]
// 3.编程，遍历出所有的monster结果，并显示在终端
// 4.提示，遍历时先取出所有keys---->keys monster*

package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Monster struct {
	name  string
	age   int
	skill string
}

func NewMonster(name string, age int, skill string) *Monster {
	return &Monster{name, age, skill}
}
func addHash(c redis.Conn, k string, m *Monster) (interface{}, error) {
	rep, err := c.Do("hmset", k, "name", m.name, "age", m.age, "skill", m.skill)
	return rep, err
}
func main() {
	// 创建3个monster
	m1 := NewMonster("qiu", 11, "sing")
	m2 := NewMonster("chen", 12, "dance")
	m3 := NewMonster("lin", 19, "swim")
	// 连接redis
	con, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("connect failed,", err)
		return
	}
	defer con.Close()
	// 将数据存放到redis中
	addHash(con, "monster1", m1)
	addHash(con, "monster2", m2)
	addHash(con, "monster3", m3)
	// 获取redis中的值
	keys, err := redis.Strings(con.Do("keys", "monster*"))
	if err != nil {
		fmt.Println("get keys failed,", err)
		return
	}
	for _, key := range keys {
		getHash(key, con)
	}
}
func getHash(key string, con redis.Conn) {
	list, err := redis.Strings(con.Do("hmget", key, "name", "age", "skill"))
	if err != nil {
		fmt.Println("hmget failed,", err)
		return
	}
	fmt.Printf("%v的name:%v, age:%v, skill:%v\n", key, list[0], list[1], list[2])
}
