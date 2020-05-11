package main

import "fmt"

//channel  基本使用

//  1.channel是引用类型
//  2.channel必须初始化才能写入数据，即make后才能使用
//  2.管道始有类型的，intChan 只能写入整型int
func main() {

	// 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	//查看channel 里面有什么
	fmt.Printf("intChan :%v intChan本身的地址%p\n", intChan, &intChan)

	// 向管道中写入数据
	intChan <- 10
	num := 32
	intChan <- num

	// 注意 ，当我们给管道写入数据时，不能超过去容量
	intChan <- 324

	// 查看 channel 的长度和容量
	fmt.Printf(" channel len=%v  cap=%v \n", len(intChan), cap(intChan))
	var num2 int
	//读取channel中的数据
	num2 = <-intChan
	fmt.Println("num2 : ", num2)

	// 查看 channel 的长度和容量
	fmt.Printf(" channel len=%v  cap=%v \n", len(intChan), cap(intChan))

	// channel 注意事项

	// 1. channel中只能存放指定的数据类型
	// 2. channel的数据放满后，就不能再存放了
	// 3. 从channel取出数据后，可以继续放入数据
	// 4. 再没有使用协程的情况下，如果channel数据取完了，再取，就会报错  dead lock

}
