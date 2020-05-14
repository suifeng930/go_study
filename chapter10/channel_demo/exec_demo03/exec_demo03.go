package main

import "fmt"

// 使用select  解决从管道取数据的阻塞问题
func main() {

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "Hello" + fmt.Sprintf("%d", i)
	}
	// 传统方法在遍历管道数据时，如果不关闭会阻塞导致dead lock
	// 使用select{} 解决 dead lock
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intCHan 读取数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从intCHan 读取数据%s\n", v)
		default:
			fmt.Printf("娶不到数据，不玩了、\n")
			return
		}
	}
}
