package main

import (
	"fmt"
	"go_study/chapter07/abstract/encapsulation"
)

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {
	//确认密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码有误")
		return
	}
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) Query(pwd string) {
	//确认密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码有误")
		return
	}

	fmt.Println("您的余额为：", account.Balance)

}
func main() {
	account := Account{
		AccountNo: "no13333",
		Pwd:       "123456",
		Balance:   124510.0,
	}

	account.Query("123456")

	// 使用 封装的特性
	newPerson := encapsulation.NewPerson("Tom")
	newPerson.SetAge(16)
	newPerson.SetSal(12312.31)

	fmt.Println("person is :", newPerson)

}

// 封装(encapsulation)  就是把抽象出来的字段和对字段的操作封装在一起， 数据被保护在内部，程序的其他包只有通过被授权的操作(方法)，才能对字段进行操作
////  封装的好处： 1 隐藏实现细节
////              2.可以对数据进行验证，保证安全合理
////  如何体现封装的特性：  1.对结构体中的属性进行封装
////                       2.通过方法、 包实现封装
//
////  实现步骤  并没有特别强调封装的概念
////          1.将结构体、字段的首字母小写
////          2.给结构体所在的包提供一个工厂模式的函数，首字母大写。 类似一个构造函数
////          3.提供一个首字母大写的Set方法(类似public) 用于对属性判断并赋值
////          4.提供一个首字母大写的Get方法(类似public) 用于获取属性的值
