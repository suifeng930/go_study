package main

import "fmt"

// 接口入门案例  高内聚  低耦合

// 创建一个接口类型， 为这个接口声明两个未被实现的方法

//  interface 类型可以定义一组方法，但这些不需要实现，并且interface不能包含任何变量，到某个自定义类型(结构体)要使用时，在根据具体情况把这些方法写出来(实现)

//  接口基本语法
//  type  interfaceName interfac{
//  method1 （paramList） return
//  method2 （paramList） return
//  }

//  注意：1. 接口里面的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想
//        2. go中的接口，不需要显示的实现，只需要一个变量，含有接口类型中的所有方法，那么这个变量就i实现了该接口，因此，golang中没有implement这样的关键字

func main() {
	// 先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点   实现接口
	computer.Working(phone)
	computer.Working(camera)

}

// 定义一个接口
type USB interface {
	// 声明两个方法
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

type Computer struct {
}

func (p Phone) Start() {
	fmt.Println(" 手机开始工作。。。")
}

func (p Phone) Stop() {
	fmt.Println(" 手机开始工作。。。")
}

// 实现usb的方法
func (p Camera) Start() {
	fmt.Println(" 相机开始工作。。。")
}

// 实现usb的方法
func (p Camera) Stop() {
	fmt.Println(" 相机开始工作。。。")
}

// 接口一个usb 接口类型
// 只要实现了Usb接口，就是指定 usb接口 声明的所有方法
func (p Computer) Working(usb USB) {
	fmt.Println(" 计算机开始工作。。。")
	usb.Start()
	usb.Stop()
}
