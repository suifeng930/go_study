package main

import "fmt"

func main() {
	var totalLevel int
	fmt.Println("请输入打印乘法表的层数")
	fmt.Scanln(&totalLevel)
	printMulti(totalLevel)
}

func printMulti(printNum int) {

	for i := 1; i <= printNum; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
