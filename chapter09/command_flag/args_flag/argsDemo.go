package main

import "fmt"
import "os"

//获取命令行的参数
//Args是一个[]string 保存了命令行的所有参数

func main() {

	fmt.Println("命令行的参数有：", len(os.Args))
	//遍历os.Args切片，就可以得到所有的命令行输入的参数值
	for key, value := range os.Args {
		fmt.Printf("args[%v]=%v \n", key, value)
	}
}
