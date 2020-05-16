package main

import (
	"fmt"
	"reflect"
)

// 反射的注意事项

//1. reflect.Value.Kind ,获取变量的类别，返回的是一个常量
//2. Type 和Kind的区别，  可能相同，也可能不同
func main() {

	var num int = 12
	Modify(&num)
	fmt.Println("num=", num)
}

//通过反射，修改 struct的值

// 通过反射的来修改变量，注意当使用setXXX() 来设置需要通过对应的指针类型来完成，这样才能改变传入的变量的值，同时需要使用到reflect.Value.Elem()
func Modify(b interface{}) {

	rVal := reflect.ValueOf(b)

	fmt.Printf("rVale =%v  type =%T", rVal, rVal)
	rVal.Elem().SetInt(123)
	fmt.Println("update rVal=", rVal)
}
