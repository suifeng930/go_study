package main

import (
	"fmt"
	"time"
)

func main() {

	chanint := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(chanint)
	go readData(chanint, exitChan)
	for {
		//遍历退出协程的函数
		ok := <-exitChan
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 0; i <= 50; i++ {
		//向channel 写数据
		intChan <- i
		fmt.Printf("writeData 写入数据 =%v\n", i)
		time.Sleep(time.Second)
	}
	//关闭channel
	close(intChan)

}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据 =%v\n", v)
		time.Sleep(time.Second)
	}
	//向 channel中插入 数据
	exitChan <- true
	close(exitChan)

}
