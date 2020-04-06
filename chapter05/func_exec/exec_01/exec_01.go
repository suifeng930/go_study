package main

import "fmt"

//  打印金字塔代码
func main() {

	var totalLevel int
	fmt.Println("请输入打印金字塔的层数")
	fmt.Scanln(&totalLevel)
	printPyramid(totalLevel)

}

func printPyramid(totalLevel int) {

	//i 表示打印的层数
	for i := 1; i <= totalLevel; i++ {
		// 在打印* 前线打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		// j 表述每层打印多少个*
		for j := 1; j <= 2*i-1; j++ {
			//if j==1|| j==2*i-1 ||i== totalLevel{   //打印空芯金字塔
			fmt.Print("*")
			//}else {
			//	fmt.Print(" ")
			//}
		}
		fmt.Println()
	}
}
