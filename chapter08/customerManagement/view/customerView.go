package main

import (
	"fmt"
	"go_study/chapter08/customerManagement/model"
	"go_study/chapter08/customerManagement/service"
)

type CustomerView struct {
	key             string                   //接受用户输入
	loop            bool                     //表示是否循环的显示主菜单
	customerService *service.CustomerService //增加一个customerService

}

func (this *CustomerView) DeleteCustomer() {
	fmt.Println("--------------------删除客户--------------------------------")
	fmt.Println("请选择你要删除的客户编码(-1)退出")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/n)")
	choice := ""
	fmt.Scanln(&choice)

	if choice == "Y" {
		//调用 删除方法
		byId := this.customerService.DeleteById(id)
		if byId {
			fmt.Println("--------删除成功--------")
		} else {
			fmt.Println("--------删除失败--------")
		}
	}

}

func (this *CustomerView) exit() {
	fmt.Println("确认退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "N" {
			break

		}
		fmt.Println("您的输入有误，请确认是否退出(Y/N):")
	}
	if this.key == "Y" {
		this.loop = false
	}

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

func (this *CustomerView) AddCustomer() {
	// 获取 当前的所有客户信息(在切片中)
	// 显示客户列表
	fmt.Println("--------------------添加客户--------------------------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomerNOID(name, gender, age, phone, email)
	if this.customerService.AddCustomer(customer) {

		fmt.Printf("--------------------添加客户成功--------------------------\n\n")
	} else {

		fmt.Printf("--------------------添加客户失败--------------------------\n\n")
	}

}

func (this *CustomerView) UpdateCustomer() {
	// 获取 当前的所有客户信息(在切片中)
	// 显示客户列表
	fmt.Println("--------------------修改客户--------------------------------")
	fmt.Println("要修改的用户编号为：")
	id := -1
	fmt.Scanln(&id)

	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别 :")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话 : ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱: ")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(id, name, gender, age, phone, email)
	if this.customerService.UpdateById(customer) {

		fmt.Printf("--------------------修改客户成功--------------------------\n\n")
	} else {

		fmt.Printf("--------------------修改客户失败--------------------------\n\n")
	}

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
			this.AddCustomer()
		case "2":
			this.UpdateCustomer()
		case "3":
			this.DeleteCustomer()
		case "4":
			this.list()
		case "5":
			this.exit()

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
