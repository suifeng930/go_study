package main

import (
	"fmt"
	"strconv"
	"time"
)

//  goroutine 入门案例

// 在主线程中，创建一个goroutine， 改协程每隔1s 输出hello world
// 在主线程中也每隔一秒 输出 hello goland  输出10次之后，退出程序
// 要求主线程和goroutine 同时执行
func main() {
	go testPrint()
	for i := 0; i < 10; i++ {
		fmt.Println("main()   hello.goland " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func testPrint() {
	for i := 0; i < 10; i++ {
		fmt.Println("testPrint()   hello.World " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//  1. 主线程是一个物理线程，直接作用在cpu上，是重量级的，非常耗费cpu资源，
//  2. 协程从主线程开启，是轻量级的线程，是逻辑态，对资源消耗相对较小
