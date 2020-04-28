package main

import (
	"fmt"
	"go_study/chapter08/customerManagement/service"
)

type CustomerView struct {
	key             string                   //接受用户输入
	loop            bool                     //表示是否循环的显示主菜单
	customerService *service.CustomerService //增加一个customerService

}

// 显示所有的客户信息
func (this *CustomerView) list() {
	// 获取 当前的所有客户信息(在切片中)
	customers := this.customerService.List()
	// 显示客户列表
	fmt.Println("--------------------客户列表--------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		info := customers[i].GetInfo()
		fmt.Println(info)

	}

	fmt.Printf("--------------------客户列表完成--------------------------\n\n")
}

func (this *CustomerView) mainMenu() {
	for {
		fmt.Println("-------------------客户信息管理软件-----------------------")
		fmt.Println("-------------------1.添加客户-----------------------")
		fmt.Println("-------------------2.修改客户-----------------------")
		fmt.Println("-------------------3.删除客户-----------------------")
		fmt.Println("-------------------4.客户列表-----------------------")
		fmt.Println("-------------------5.退出-----------------------")
		fmt.Print("请选择(1~5):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			this.list()
		case "5":
			this.loop = false

		}
		if !this.loop {
			break

		}
	}
	fmt.Println("您退出了客户管理系统")

}

func main() {
	customerView := CustomerView{
		key:  "",
		loop: true,
	}
	// 完成对customerView和customerService的方法
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()

}
