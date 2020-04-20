package main

import (
	"fmt"
)

// golang 的方法是作用在指定的数据类型上的 (和指定类型绑定)， 因此自定义类型，都可以有方法，而不仅仅是struct

func main() {

	var t Person // 声明变量
	t.Name = "Tom"
	t.test()  //调用a 的方法
	t.speak() //打印speak()
	t.jisuan()
	t.countNum(13)
	t.countValue(13, 24)
}

//方法的声明和调用
type Person struct {
	Name string
}

// 给A 绑定一个方法
// test（） 只能通过Person类型的变量调用，不能通过其他类型变量调用
func (a Person) test() {
	fmt.Println("test() ", a.Name)
}

func (person Person) speak() {
	fmt.Printf("speak()  %s is a good man \n", person.Name)
}

// 给person 类型绑定一个方法
func (person Person) jisuan() {
	res := 0
	for i := 1; i < 200; i++ {
		res += i
	}
	fmt.Printf("jisuan() %s count value is %d\n", person.Name, res)
}

//给person 接口一个数n,计算 1+n的结果
func (person Person) countNum(intNum int) {
	res := 0
	for i := 1; i < intNum; i++ {
		res += i
	}
	fmt.Printf("jisuan() %s count value is %d\n", person.Name, res)
}

////给person 接口一个数n,计算 1+n的结果
func (person Person) countValue(intNum int, res int) int {
	for i := 1; i < intNum; i++ {
		res += i
	}
	fmt.Printf("jisuan() %s count value is %d\n", person.Name, res)
	return res
}
