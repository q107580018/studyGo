package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Connect redis failed,", err)
		return
	}
	defer c.Close()
	rep, err := c.Do("set", "name1", "qiuwen") // 写入string数据

	if err != nil {
		fmt.Println("redis set failed,", err)
		return
	}
	fmt.Println(rep)
	rep, err = redis.String(c.Do("get", "name1"))
	if err != nil {
		fmt.Println("redis Get failed,", err)
	}
	fmt.Println(rep)
}
