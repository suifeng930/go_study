package main

import (
	"fmt"
	"time"
)

//错误处理机制

//1.  go语言追求简洁优雅，所以go语言不支持传统的 异常处理机制
//2. go中引入的处理方式为： defer    panic      recover

//3.  这几个异常的使用场景可以这么简单描述， go 中可以抛出一个panic的异常，然后在defer中通过 recover捕获该异常，然后正常处理
func main() {
	test()
	for {
		fmt.Println("继续执行下面的代码")

		time.Sleep(time.Second)
	}

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
