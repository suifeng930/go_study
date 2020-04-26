package main

import "fmt"

// 家庭记账转件

// 掌握技能： 编码技巧+调试技巧
//           1.局部变量和基本数据类型
//           2.循环语句
//           3.分支语句
//           4.简单的屏幕输出格式控制
//           5.进阶：面向对象编程使用

// 需求说明：
//           1.模拟实现基于文本界面的<家庭记账软件>
//           2.该软件能够记录家庭的收入、支出，并能够打印收支明细表
//           3.项目采用分级菜单方式。主菜单
//           ---------------家庭收支记账软件---------------------
//           ---------------1.收支明细---------------------
//           ---------------2.登记收入---------------------
//           ---------------3.登记支出---------------------
//           ---------------4.退出---------------------
//           ---------------请选择（1~4）---------------------

//          把实现记账 显示明细放到一个FamilyAccount结构体中 oop编程模式
func main() {
	account := NewFamilyAccount()
	username := ""
	password := ""
	for i := 0; i <= 3; i++ {
		fmt.Println("请输入你的用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if username == "xiaoma" && password == "123456" {
			fmt.Printf("欢迎 %s你使用，该计算软件\n", username)
			account.MainMenu()
			break
		} else {
			fmt.Println("\n密码或用户名输入错误，请重新输入！")
			fmt.Printf("\n第%d次输入密码或用户名错误，请重新输入！你还有第%d次机会", i, 3-1)
			if i == 3 {
				fmt.Println("对不起，你输入密码的次数已经超过了三次，系统将锁定")

			}
		}
	}

}

type FamilyAccount struct {
	balance float64 //账户余额
	money   float64 //金额
	note    string  // 收支明细
	details string  // 收支明细
	key     string
	loop    bool // 声明一个字段，控制是否退出for
	flag    bool // 记录是否存在收支记录
}

//编写一个工厂模式的狗找方法，返回一个 *FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		balance: 12323.0,
		money:   0,
		note:    "",
		details: "收支\t 余额 \t金额 \t说明",
		key:     "",
		loop:    true,
		flag:    false,
	}
}

func (this *FamilyAccount) MainMenu() {

	for {
		fmt.Println("\n---------------家庭收支记账软件---------------------")
		fmt.Println("---------------1.收支明细---------------------")
		fmt.Println("---------------2.登记收入---------------------")
		fmt.Println("---------------3.登记支出---------------------")
		fmt.Println("---------------4.退出-------------------------")
		fmt.Println("---------------4.退出请选择（1~4）---------------------")
		fmt.Print("请选择(1~4):")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.Income()
		case "3":
			this.Pay()
		case "4":
			this.Exit()
		default:
			fmt.Println("请输入正确的选项。。。。")
		}
		if !this.loop {
			break //跳出循环，退出
		}
	}
	fmt.Println("你退出了家庭记账软件使用")
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("---------------收支明细---------------------")
	if this.flag {
		fmt.Println(this.details)

	} else {
		fmt.Println("当前没有收入哟，来一笔吧")
	}
}
func (this *FamilyAccount) Income() {
	fmt.Println("---------------登记收入---------------------")
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入 \t%v \t%v \t%v ", this.balance, this.money, this.note)
	this.flag = true

}

func (this *FamilyAccount) Pay() {
	fmt.Println("---------------登记支出---------------------")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("支出余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次收支说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出 \t%v \t%v \t%v ", this.balance, this.money, this.note)
	this.flag = true

}

func (this *FamilyAccount) Exit() {
	fmt.Println("你确定要退出嘛？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if choice == "y" {
		this.loop = false
	}

}
