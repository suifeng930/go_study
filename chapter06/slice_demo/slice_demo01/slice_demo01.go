package main

import "fmt"

//   切片的基本概念
//  1. 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递规则
//  2. 切片的使用和数组类似，遍历切片，访问切片的元素和切片长度len 都一样
//  3. 切片的长度是可以变化的，因此切片是一个动态变化的数组
//  4. 切片定义语法    var sliceName []类型

func main() {

	var intStr [5]int = [...]int{1, 234, 544, 13, 5}

	//切片的基本使用
	slice := intStr[1:3]
	fmt.Println("intStr=", intStr)
	fmt.Println("slice的元素是： ", slice)
	fmt.Println("slice的元素的个数： ", len(slice))
	fmt.Println("slice的元素的容量： ", cap(slice))
	fmt.Printf("intStr[1]元素的地址是：%p \n", &intStr[1])
	fmt.Printf("slice[0]元素的地址是：%p \n", &slice[0])

	// 通过make 创建切片    基础语法： var  sliceName []type = make([]type,len,[cap])
	//
	intSlice := make([]int, 3, 4)

	for i := 0; i < len(intSlice); i++ {
		intSlice[i] = i + 10
	}
	fmt.Println("intSlice value :", intSlice)
	fmt.Println("intSlice 的caps :", cap(intSlice))

	// 定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var StrSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("StrSlice=", StrSlice)
	fmt.Println("StrSlice size = ", len(StrSlice))
	fmt.Println("StrSlice cap = ", cap(StrSlice))

	// 遍历切片
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("intSlice [%v]=%v\n", i, intSlice[i])
	}
	for i, v := range intSlice {
		fmt.Printf("intSlice [%v]=%v\n", i, v)
	}

	// cap是一个内置函数吗，用于统计切片的容量，即最大可以存放多少个元素
	// 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组中，或者使用make定义

}
