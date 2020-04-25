package main

import "fmt"

// 接口的应用场景

//  接口的注意事项和细节
//  1. 接口本省不能创建实例，但是可以指向一个实现了该接口类型的自定义类型的变量(实例)
//  2. 接口中所有的方法都没有方法体，即都是没有实现的方法
//  3. 在go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
//  4. 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
//  5. 一个自定义类型可以实现多个接口(多继承)
//  6. go接口中不能有任何变量
//  7. 一个接口(A 接口) 可以继承多个别的接口(B、C接口) 这时如果要实现A接口，也必须全部实现B\C 接口的方法
//  8. interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用那么会输出nil
//  9. 空接口interface没有任何方法，所以所有类型都实现了空接口

func main() {
	var a AInterface //初始化一个空接口
	student := Student{}
	student.Name = "hello"
	a = student //接口必须指向一个实现它的结构体变量
	a.Say()

	fmt.Println("golang中一个结构体可以实现多个接口")
	var monstar Monstar
	var a1 AInterface = monstar
	var b1 BInterface = monstar
	a1.Say()
	b1.Hello()

}

type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("stu  Say  student")
}

type Monstar struct {
	Name string
}

func (m Monstar) Hello() {
	fmt.Println(" monstar implement  BInterface  ...")

}

func (m Monstar) Say() {
	fmt.Println(" monstar implement  AInterface  ...")

}
