package main

import (
	"errors"
	"fmt"
)

//错误处理机制

//1.  go语言追求简洁优雅，所以go语言不支持传统的 异常处理机制
//2. go中引入的处理方式为： defer    panic      recover

//3.  这几个异常的使用场景可以这么简单描述， go 中可以抛出一个panic的异常，然后在defer中通过 recover捕获该异常，然后正常处理
func main() {
	//test()
	//for {
	//	fmt.Println("继续执行下面的代码")
	//
	//	time.Sleep(time.Second)
	//}

	//自定义测试
	test02()

}

func test() {

	// 使用defer + recover 来捕获和处理异常

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err= ", err)
			// 记录错误日志之后，可以加入一些告警日志，等操作
			fmt.Println("发送邮件给admin user ")
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("Err=", res)
}

// 自定义错误
// go  支持自定义错误，使用error.new 和panic内置函数
// errors.new("错误说明“)  ，会返回一个error类型的值，表示一个错误
// panic内置函数，接收一个interface{} 类型的值作为参数，可以接受error类型的变量，输出错误信息，并退出程序

// 函数去读取配置文件中的init.conf的信息
// 如果文件名称传入不正确，我么就返回一个自定义的错误

func readConf(fileName string) (err error) {

	if fileName == "config.ii" {
		//读取正确
		return nil
	} else {
		// 返回一个自定义错误类型
		return errors.New("读取文件失败。。。")

	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {

		// 如果读取文件失败，输出错误，并终止程序
		panic(err)
	}
	fmt.Println("test02() 继续执行。。。")
}
