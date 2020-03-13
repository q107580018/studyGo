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
	// 通过go 向redis写入hash数据
	_, err = c.Do("hset", "user", "name", "qiuwen", "age", 18, "sex", "male")
	if err != nil {
		fmt.Println("Hset failed,", err)
		return
	}
	// 通过go 向redis读取hash数据
	for _, key := range []string{"name", "age", "sex"} {
		rep, err := redis.String(c.Do("hget", "user", key))
		if err != nil {
			fmt.Println("Hget failed,", err)
			return
		}
		fmt.Printf("%v:%v\n", key, rep)
	}
}
