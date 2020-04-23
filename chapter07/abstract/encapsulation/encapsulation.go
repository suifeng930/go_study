package encapsulation

import "fmt"

// 封装(encapsulation)  就是把抽象出来的字段和对字段的操作封装在一起， 数据被保护在内部，程序的其他包只有通过被授权的操作(方法)，才能对字段进行操作
//  封装的好处： 1 隐藏实现细节
//              2.可以对数据进行验证，保证安全合理
//  如何体现封装的特性：  1.对结构体中的属性进行封装
//                       2.通过方法、 包实现封装

//  实现步骤  并没有特别强调封装的概念
//          1.将结构体、字段的首字母小写
//          2.给结构体所在的包提供一个工厂模式的函数，首字母大写。 类似一个构造函数
//          3.提供一个首字母大写的Set方法(类似public) 用于对属性判断并赋值
//          4.提供一个首字母大写的Get方法(类似public) 用于获取属性的值

// 要求创建一个person ，不能随便查看人的年龄 工资等隐私，并对输入的年龄进行合理的验证

type person struct {
	Name string

	age int
	sal float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
		age:  0,
		sal:  0,
	}

}

// 提供set方法
func (ps *person) SetAge(age int) {
	if age > 0 && age < 120 {
		ps.age = age
	} else {
		fmt.Println("输入的age  不正确")
	}

}

func (ps *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		ps.sal = sal
	} else {
		fmt.Println("输入的sal  不正确")
	}

}

// 提供get方法
func (ps *person) GetSal() float64 {
	return ps.sal

}
func (ps *person) GetAge() int {
	return ps.age

}
