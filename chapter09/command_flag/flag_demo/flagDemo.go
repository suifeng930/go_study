package main

import (
	"flag"
	"fmt"
)

// flag包的基本使用
func main() {

	// 定义几个变量用于接受命令行的参数值

	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "localhost", "host")
	flag.IntVar(&port, "p", 0, "port")
	flag.Parse() // 讲数据转换
	fmt.Printf("user=%v ,pwd=%v ,host=%v ,post=%d", user, pwd, host, port)
}
