package main

import (
	"fmt"
	"go_study/chapter11/chatRoom/client/process"
)

var userId int
var userKey, userName string

func main() {

	var key int // 接受用户的选择
	//var loop = true //判断是否继续显示菜单

	for true {
		fmt.Println("----------------欢迎登录多人聊天系统----------------------")
		fmt.Println("\t\t\t1. 登录聊天室")
		fmt.Println("\t\t\t2. 注册用户")
		fmt.Println("\t\t\t3. 退出系统")
		fmt.Println("\t\t\t 请选择(1~3)")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入您的 ID好")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userKey)

			//完成登录操作，创建一个userProcess()
			userProcess := &process.UserProcess{}
			userProcess.Login(userId, userKey)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userKey)
			fmt.Println("请输入用户名")
			fmt.Scanf("%s\n", &userName)
			// 调用User Process 完成用户注册的请求
			userProcess := &process.UserProcess{}
			userProcess.Register(userId, userKey, userName)

		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有误，请重新输入")

		}
	}
	////根据用户的输入，显示新的提示信息
	//if key == 1 {
	//	//用户要登录了
	//	fmt.Println("请输入您的 ID好")
	//	fmt.Scanf("%d\n", &userId)
	//	fmt.Println("请输入用户的密码")
	//	fmt.Scanf("%s\n", &userKey)
	//	//先把登录的用户函数，写到另外一个文件中 login.go
	//	err := login.Login(userId, userKey)
	//	if err != nil {
	//		fmt.Println("登录失败")
	//	} else {
	//		fmt.Println("登录成功")
	//	}
	//
	//} else if key == 2 {
	//	fmt.Println("注册用户")
	//
	//} else {
	//
	//}

}
