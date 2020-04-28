package service

import "go_study/chapter08/customerManagement/model"

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
