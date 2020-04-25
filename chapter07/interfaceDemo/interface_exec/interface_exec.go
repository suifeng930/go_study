package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//  当A结构体继承了B结构体，那么A结构体就自动继承了B结构体的字段和方法，并且可以直接使用
//  当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此我们认为，实现接口是对继承机制的补充

//   实现接口可以看作是对 继承的一种补充
//   接口和继承解决的问题是不同的
//      继承的价值： 解决代码的复用性和可维护性
//      接口的价值： 设计，设计好各种规范(方法)，让其他他自定义类型去实现这些方法
//      接口比继承更加灵活： 继承是满足 is -a 的关系， 但是接口时满足 like-a的关系
//      接口在一定程度上实现代码的解耦
//      多态的特性：        多态特性是通过接口实现的，可以按照统一的接口来调用不同的实现，这时接口变量就呈现不同的形态

//      多态参数：   一个接口，接收不同的参数变量，来体现多态的特性： 一个USB接口，既可以接受手机变量，也可以接受相机变量，此时这个USB接口就体现了多态
//      多态数组：   给 USB数组种，存放Phone结构体和Camera 结构体变量，Phone还有一个特有的方法call（）  请遍历USB数组，如果时phone变量，除了调用USB接口声明的方法外，还需要调用Phone特有的方法Call()

func main() {

	// 先定义一个数组/切片
	var intSlice = []int{0, -1, 21, 43, 6}
	// 要求对intSlice切片进行排序
	// 1. 冒泡排序
	// 2. 使用系统提供的排序算法
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	// 如何对结构体进行排序
	// 1.冒泡排序
	// 2.sort.Sort（）
	var Heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		Heros = append(Heros, hero)
	}
	fmt.Println("Sort Before", Heros)
	// 调用排序的方法
	sort.Sort(Heros)
	fmt.Println("Sort After :", Heros)

}

// 1.声明一个Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2. 声明一个hero结构体的切片类型
type HeroSlice []Hero

// 3.实现interface接口
//  获取slice数组长度
func (hs HeroSlice) Len() int {
	return len(hs)

}

// 决定排序方法规则
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age

}

//交换数据
func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp

}

// 接口体现多态的  多态数组形式
