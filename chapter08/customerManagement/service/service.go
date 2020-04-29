package service

import (
	"go_study/chapter08/customerManagement/model"
)

//完成对customer的操作，包括crud
type CustomerService struct {
	custmers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段还可以作为新客户的id
	customerNum int
}

// 编写一个方法 返回一个 customerService 指针
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "123@qq.com")
	customerService.custmers = append(customerService.custmers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.custmers

}

// 新增一个客户信息
func (this *CustomerService) AddCustomer(customer model.Customer) bool {

	this.customerNum += 1
	customer.Id = this.customerNum
	this.custmers = append(this.custmers, customer)

	return true

}

//根据id 查询客户在切片中对饮下标，如果没有该客户返回 -1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for key := range this.custmers {
		if key == id {
			index = key
		}
	}
	return index
}

//根据id 删除用户
func (this *CustomerService) DeleteById(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false

	}
	// 如何从求篇中删除一个元素  切片的特性
	this.custmers = append(this.custmers[:index], this.custmers[index+1:]...)
	return true

}

//根据id 删除用户
func (this *CustomerService) UpdateById(customer model.Customer) bool {
	index := this.FindById(customer.Id)
	if index == -1 {
		return false

	}
	this.custmers[index] = customer
	return true

}
