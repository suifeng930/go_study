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
	fmt.Printf("totalWeight=%v  avgWeight=%v", totalWeight, avgWeight)
}
