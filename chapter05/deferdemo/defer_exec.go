package main

import (
	"fmt"
	"log"
)

// defer 的主要价值在，当函数执行完毕后，可以及时的释放函数创建的资源
// defer 在创建资源(比如 ：数据库连接、文件句柄、锁等）,为了函数执行完毕之后，即使的释放资源，go 提供了defer延时调用机制
func main() {

	result := sum(10, 20)

	log.Println(" result=", result)
}

func sum(num1, num2 int) int {

	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的堆栈中(defer栈)
	// 当函数执行完毕后，再从defer栈中 ，按照先进后出的方式出栈，执行
	defer fmt.Println(" ok1 num1=", num1) // ok1 num1=10
	defer fmt.Println(" ok2 num2=", num2) // ok2 num2=20

	res := num1 + num2            // res= 30
	fmt.Println(" ok3 res=", res) //1.  ok3  res=30
	return res
}
