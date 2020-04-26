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

//          功能1.  先完成可以显示主菜单，并且可以退出 (当用户输出4 就退出)
func main() {

	key := ""
	flag := true
	// 收支余额
	balance := 12312.00
	//收支金额
	money := 0.0
	//详细说明
	note := ""
	// details 详细说明|
	details := "收支\t 余额 \t金额 \t说明"

	//定义一个变量 记录是否有收支行为
	FLAG := false

	for {
		fmt.Println("\n---------------家庭收支记账软件---------------------")
		fmt.Println("---------------1.收支明细---------------------")
		fmt.Println("---------------2.登记收入---------------------")
		fmt.Println("---------------3.登记支出---------------------")
		fmt.Println("---------------4.退出-------------------------")
		fmt.Println("---------------4.退出请选择（1~4）---------------------")
		fmt.Print("请选择(1~4):")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("---------------收支明细---------------------")
			if FLAG {
				fmt.Println(details)

			} else {
				fmt.Println("当前没有收入哟，来一笔吧")
			}

		case "2":
			fmt.Println("---------------登记收入---------------------")
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入 \t%v \t%v \t%v ", balance, money, note)
			FLAG = true
		case "3":
			fmt.Println("---------------登记支出---------------------")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("支出余额不足")
				break
			}
			balance -= money
			fmt.Println("本次收支说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出 \t%v \t%v \t%v ", balance, money, note)
			FLAG = true
		case "4":
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
				flag = false
			}
		default:
			fmt.Println("请输入正确的选项。。。。")
		}
		if !flag {
			break //跳出循环，退出
		}
	}
	fmt.Println("你退出了家庭记账软件使用")
}
