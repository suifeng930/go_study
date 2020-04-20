package main

import "fmt"

//方法的调用和传参机制原理：
// 1.方法的调用和传参机制和函数基本一致，但是方法调用时，会将调用方法的变量，当作实参也传递给方法
//2.变量调用方法时，变量本身也会作为一个参数传递到方法中，可以是值类型，可以是引用类型

func main() {
	var c Circle
	c.Radius = 12.4
	c.Area() //

}

// 编写一个结构体 Circle, 字段为radius
// 声明一个方法 Area 和Circle绑定，可以返回面积
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	pia := 3.14
	area := pia * c.Radius * c.Radius
	fmt.Println("area() the area is ", area)
	return area
}

//为了提高效率，通常将方法和struct的指针类型绑定
func (c *Circle) Areas() float64 {

	pia := 3.14
	area := pia * c.Radius * c.Radius
	fmt.Println("area() the area is ", area)
	return area
}

// func （receiver type） methodName(参数列表)（返回值列表）{
// 方法体
// 返回值
// }
