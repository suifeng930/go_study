package main

import (
	"fmt"
	"reflect"
)

// 反射快速入门
//  演示对基本数据类型，interface   reflect.Value()进行反射的基本操作
func main() {

	var num int
	reflectTest01(num)

}

type Student struct {
	Name string
	Age  int
}

func reflectTest01(b interface{}) {
	// 通过反射获取传入的变量的type, kind,
	//1. 先获取reflect.Type

	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	// 2. 获取到reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue =%d rrValue type= %T\n", rValue, rValue)

	//3/下面将 rValue  转成interface()
	Interface := rValue.Interface()
	num := Interface.(int)
	fmt.Println("num =", num)

}
