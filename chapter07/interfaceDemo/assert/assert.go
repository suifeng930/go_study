package main

import (
	"fmt"
)

// 类型断言

//  2.需求： 如何将一个接口变量，赋给自定义类型的变量--->引出类型断言
//   类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
func main() {

	var a interface{}
	var point Point = Point{
		x: 1,
		y: 2,
	}
	a = point // 将point  赋值给一个空接口

	var b Point
	//  b =a   会报错，如何解决
	//这里使用断言
	//   类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
	b = a.(Point) // 断言A  的类型是不是Point类型
	fmt.Println(b)

	fmt.Println("类型断言的其他案例")
	var xs interface{}
	var b2 float32 = 13.3
	xs = b2 // 空接口可以接受任意类型
	Y, ok := xs.(float32)
	if ok {
		fmt.Println("convert success")
		fmt.Printf(" y 的类型是 %T  值是%v\n", Y, Y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("在进行类型断言是，如果类型不匹配，就会报错panic,因此在进行类型断言时，要确保原来的空接口指向的就是要断言的类型")
	// 2.如何在进行断言时，带上检测机制，如果成功ok,否则也不要报错panic

}

type Point struct {
	x int
	y int
}
