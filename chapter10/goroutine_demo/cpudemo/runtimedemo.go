package main

import (
	"fmt"
	"runtime"
)

//获取当前主机的cpu核心数
func main() {

	// 获取当前系统cpu的核心数
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("cpuNum :", numCPU)

}
