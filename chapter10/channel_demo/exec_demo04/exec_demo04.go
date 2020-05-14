package main

import (
	"log"
	"time"
)

// goroutine 中使用 recover，可以解决协程中出现panic, 导致程序崩溃问题

func main() {
	go PrintHello()
	go Test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.Println("main() running")

	}
	// 如果我们起了一个协程，但这个协程出现了panic  如果我们没有捕获这个panic,就会造成整个程序崩溃，这时，我们可以在goroutine中使用recover+defer来捕获panic，
	// 这样及时一个协程出现了问题，也不会导致整个程序的不可用

}

func PrintHello() {
	for i := 0; i < 10; i++ {
		log.Println("Hello world")
		time.Sleep(time.Second)

	}
}

func Test() {

	//使用 defer +recover解决
	defer func() {
		// 捕获 test() 的panic
		if err := recover(); err != nil {
			log.Println("test() 发生错误了，", err)

		}
	}()
	//定义一个map
	var myMap map[int]string

	myMap[0] = "goland" //error

}
