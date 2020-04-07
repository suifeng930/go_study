package main

import "fmt"

func main() {

	num1 := 100
	fmt.Printf("num1的类型%T, num1的值为 =%v num1的地址为=%v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的类型%T, num2的值为 =%v num2的地址为=%v\n", num2, num2, &num2)

}
