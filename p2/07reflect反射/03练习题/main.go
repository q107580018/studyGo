package main

import (
	"fmt"
	"reflect"
)

// 给你一个变量 var v float64=1.2， 请使用反射来得到他的reflect.Value
// 然后获取对应的Type,kind和值,并将reflect.Value转换成interface{},再将
// interface{}转换成float64.
func main() {
	var v float64 = 1.2
	rVal := reflect.ValueOf(v)
	rType := reflect.TypeOf(v)
	IV := rVal.Interface()
	rFloat := IV.(float64)
	fmt.Printf("rValue:%v, rType:%v, rFloat:%v\n", rVal, rType, rFloat)
}
