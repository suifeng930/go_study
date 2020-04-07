package main

import (
	"fmt"
	"strconv"
	"time"
)

// 时间和日期相关函数学习

func main() {

	//1. 获取当前时间
	newTime := time.Now()
	fmt.Printf("type  =%T var=%v \n", newTime, newTime)

	//2. 获取当前时间的年月日，时分秒
	fmt.Printf("year=%v\n", newTime.Year())
	fmt.Printf("month=%v\n", newTime.Month())
	fmt.Printf("dey=%v\n", newTime.Day())
	fmt.Printf("hour=%v\n", newTime.Hour())
	fmt.Printf("Minute=%v\n", newTime.Minute())
	fmt.Printf("Second=%v\n", newTime.Second())

	//3. 格式化日期时间
	//    使用fmt.prrint()
	fmt.Printf("当前日期： %d-%d-%d %d:%d:%d \n", newTime.Year(), newTime.Month(), newTime.Day(), newTime.Hour(), newTime.Minute(), newTime.Second())
	dataTime := fmt.Sprintf("当前日期： %d-%d-%d %d:%d:%d \n", newTime.Year(), newTime.Month(), newTime.Day(), newTime.Hour(), newTime.Minute(), newTime.Second())
	fmt.Println("nowTime", dataTime)
	//    使用printFormat
	formatTime := newTime.Format("2006/01/02 15:04:05")
	fmt.Println("formatTime=", formatTime)

	//4. 时间休眠  睡眠一秒
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	// 获取当前unix 时间戳，和unixnano 时间错。 作用获取生成随机数

	fmt.Printf("unix 时间戳=%v  unixnano时间戳=%v \n", newTime.Unix(), newTime.UnixNano())

	//获取unix时间戳
	start := time.Now().Unix()
	test02()
	end := time.Now().Unix()
	fmt.Printf("执行耗费时间=%d \n", end-start)

}

// 统计一个函数的执行时间

func test02() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
