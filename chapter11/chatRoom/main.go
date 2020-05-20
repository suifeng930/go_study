package main

import (
	"fmt"
	"go_study/chapter11/chatRoom/client"
)

var userId int
var userKey string

func main() {

	var key int     // 接受用户的选择
	var loop = true //判断是否继续显示菜单

	for loop {
		fmt.Println("----------------欢迎登录多人聊天系统----------------------")
		fmt.Println("\t\t\t1. 登录聊天室")
		fmt.Println("\t\t\t2. 注册用户")
		fmt.Println("\t\t\t3. 退出系统")
		fmt.Println("\t\t\t 请选择(1~3)")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")

		}
	}
	//根据用户的输入，显示新的提示信息
	if key == 1 {
		//用户要登录了
		fmt.Println("请输入您的 ID好")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", userKey)
		//先把登录的用户函数，写到另外一个文件中 login.go
		err := client.Login(userId, userKey)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}

	} else if key == 2 {
		fmt.Println("注册用户")

	} else {

	}

}
