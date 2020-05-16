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

	fmt.Println("======================反射 struct ================")
	student := Student{
		Name: "tom",
		Age:  123,
	}
	reflect02(student)

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

//反射用例 对struct的用法
func reflect02(b interface{}) {
	// 1. 先获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	//2. 获取到reflect.value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue =%v rrValue type= %T\n", rValue, rValue)
	// 将value 转化interface（）
	iV := rValue.Interface()
	fmt.Printf("iv =%v iV type = %T\n", iV, iV)
	// 获取到struct的类型，需要使用类型断言机制
	student, ok := iV.(Student)
	if ok {
		fmt.Println("使用类型断言获取，student=", student.Name)
	}
}
