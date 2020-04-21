package main

import "fmt"

//方法的调用和传参机制原理：
// 1.方法的调用和传参机制和函数基本一致，但是方法调用时，会将调用方法的变量，当作实参也传递给方法
//2.变量调用方法时，变量本身也会作为一个参数传递到方法中，可以是值类型，可以是引用类型

func main() {
	var c Circle
	c.Radius = 12.4

	c.Area() //

	var method MethodUtils
	method.Print()
	area := method.PrintArea(113, 13)
	fmt.Println(area)
	fmt.Println()
	var cal Calcuator
	cal.Num1 = 123.2
	cal.Num2 = 143.2
	res := cal.getRes('+')
	fmt.Println(res)
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

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) PrintArea(len int64, width int64) int64 {

	return len * width
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误。。。")

	}
	return res

}

// 方法和函数的区别
//1. 调用方式不一样
//   函数调用方式：   函数名( 实参列表)
//   方法调用方式：   变量.方法名( 实参列表)
//2. 对于普通函数，接收者为值类型的时候，不能将指针类型的数据直接传递，反之亦然
//3. 对于方法(如 struct的方法)， 接收者为值类型时，可以直接用指针类型变量调用方法，反之亦然
