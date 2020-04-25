package main

import "fmt"

// 多态数组

//      多态的特性：        多态特性是通过接口实现的，可以按照统一的接口来调用不同的实现，这时接口变量就呈现不同的形态

//      多态参数：   一个接口，接收不同的参数变量，来体现多态的特性： 一个USB接口，既可以接受手机变量，也可以接受相机变量，此时这个USB接口就体现了多态
//      多态数组：   给 USB数组种，存放Phone结构体和Camera 结构体变量，Phone还有一个特有的方法call（）  请遍历USB数组，如果时phone变量，除了调用USB接口声明的方法外，还需要调用Phone特有的方法Call()

func main() {
	// 定义一个接口数组  ，可以存放 phone he camera 的结构体变量
	var UsbArr [3]USB
	UsbArr[0] = Phone{"XIAOMI"}
	UsbArr[1] = Camera{"SONY"}
	UsbArr[2] = Phone{"APPLE"}

	// 打印多态数组
	fmt.Println("打印多态数组", UsbArr)
	var computer Computer
	for key := range UsbArr {
		fmt.Printf("--------%v---------\n", UsbArr[key])
		computer.Working(UsbArr[key])

	}
}

// 定义一个接口
type USB interface {
	// 声明两个方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

func (p Phone) Start() {
	fmt.Println(" 手机开始工作。。。")
}

// 手机的自有功能  打电话
func (p Phone) Call() {
	fmt.Println(p.Name, " 手机正在打电话。。。")
}
func (p Phone) Stop() {
	fmt.Println(" 手机开始工作。。。")
}

// 实现usb的方法
func (p Camera) Start() {
	fmt.Println(" 相机开始工作。。。")
}

// 实现usb的方法
func (p Camera) Stop() {
	fmt.Println(" 相机开始工作。。。")
}

// 接口一个usb 接口类型
// 只要实现了Usb接口，就是指定 usb接口 声明的所有方法
func (p Computer) Working(usb USB) {
	fmt.Println(" 计算机开始工作。。。")
	usb.Start()
	usb.Stop()
	//  使用类型断言，判断是不是手机
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
}
