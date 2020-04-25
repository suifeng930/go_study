package main

import "fmt"

// 编写一个函数，可以判断输入的参数是什么类型

func main() {

	var ne float64 = 2.3
	var f float32 = 1.2
	var n int32 = 30
	address := "北京"

	//类型判断
	TypeJudge(ne, f, n, address)
}

func TypeJudge(items ...interface{}) {

	for key, value := range items {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool 类型： 值是%v \n", key, value)
		case float32:
			fmt.Printf("第%v个参数是float32 类型： 值是%v \n", key, value)
		case float64:
			fmt.Printf("第%v个参数是float64 类型： 值是%v \n", key, value)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int 类型： 值是%v \n", key, value)
		case string:
			fmt.Printf("第%v个参数是string 类型： 值是%v \n", key, value)
		default:
			fmt.Printf("第%v个参数 类型不确定： 值是%v \n", key, value)

		}

	}
}
