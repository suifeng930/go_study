package main

import (
	"fmt"
	"reflect"
)

// 反射实践
//1. 使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
func main() {

	var a Monster = Monster{
		Name:  "孙悟空",
		Age:   4540,
		Skill: "金箍棒",
	}
	TestStruct(a)

}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Skill string
}

func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")

}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, skill string) {
	s.Name = name
	s.Age = age
	s.Skill = skill
}

func TestStruct(b interface{}) {

	val := reflect.ValueOf(b)
	typeOf := reflect.TypeOf(b)
	// 获取类别
	kind := val.Kind()
	if kind != reflect.Struct { // 判断是不是传入的结构体
		fmt.Println("expect struct")
		return
	}
	// 获取到结构体的字段总数
	numField := val.NumField()
	fmt.Printf("struct has %d fields\n", numField)
	for i := 0; i < numField; i++ {
		fmt.Printf("Field %d  value=%v\n", i, val.Field(i))

		//获取到struct的标签，注意需要通过reflect.Type 来获取tag把标签的值
		tagVal := typeOf.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d ; tag=%v\n", i, tagVal)

		}
	}
	//h获取结构体的方法总数
	numMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numMethod)
	//
	val.Method(1).Call(nil)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res =", res[0].Int())

}
