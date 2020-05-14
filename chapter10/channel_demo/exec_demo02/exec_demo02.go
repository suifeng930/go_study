package main

import (
	"log"
)

// 统计1~20000的数字中，哪些是素数，采用 goroutine  + channel完成
func main() {

	chanInt := make(chan int, 1000)
	peimeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4) //退出主线程

	go putNum(chanInt)

	for i := 0; i < 4; i++ {
		go primeNum(chanInt, peimeChan, exitChan)
	}

	// 主线程 要在线程中获取到exitChan
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当从这个管道中取出了4个结果
		close(peimeChan)
	}()

	//遍历peimeChan  把结果取出
	for {
		res, ok := <-peimeChan
		if !ok {
			break
		}

		log.Printf("素数=%d\n", res)
	}
	log.Println("主线程退出。。。")
}

// 往chan中写入数据
func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

// 判断一个数是不是素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		//判断num是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明不死素数
				flag = false
				break
			}
		}
		if flag {
			// num为素数
			primeChan <- num
		}
	}
	log.Println("有一个primeChan  协程因为取不到数据， 退出")
	exitChan <- true
}
