package main

import "fmt"

//  在go 中  数组是值类型
func main() {

	// 初始化一个数组
	var hens [6]float64
	//为数组赋值
	hens[0] = 3.8
	hens[1] = 3.1
	hens[2] = 3.4
	hens[3] = 3.2
	hens[4] = 3.7
	hens[5] = 3.6
	//定义变量接受 总量
	//定义变量接受 avg
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	avgWeight := totalWeight / 6
	fmt.Printf("totalWeight=%v  avgWeight=%v\n", totalWeight, avgWeight)
	//1，数组是多个相同类型数据的组合，一个数组一旦声明/定义了其长度是固定的，不能动态变化。
	//2. 数组创建后，如果没有赋值，有默认值(零值)
	//   数值型(整数系列 浮点数系列)   0
	//3. go 的数组是属于值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会互相影响

	arr := [3]int{11, 22, 44}
	fmt.Println(arr)
	test01(arr)
	fmt.Printf("arr 的地址是=%p \n", &arr)
	fmt.Println("调用函数指针")
	test02(&arr)
	fmt.Println(arr)
}

// 定义一个函数，接受一个数组
func test01(arr [3]int) {

	arr[0] = 88
	fmt.Printf("test01（） arr 的地址是=%p \n", &arr)
}

// 定义一个函数，接受一个指针
func test02(arr *[3]int) {
	(*arr)[0] = 455 //通过指针去数据
	fmt.Printf("test02（） arr 的地址是=%p \n", &arr)
}
