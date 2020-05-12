package main

import (
	"fmt"
	"math/rand"
	"time"
)

//  完成如下案例
//  1.创建一个Person 结构体{Name,Age ,Address}
//  1.使用rand() 配合随机创建10个person实例，并放入到channel中
//  1.遍历channel， 将各个perSon实例打印显示再终端
func main() {

	var chanPeople chan Person
	chanPeople = make(chan Person, 11)
	for i := 0; i < 10; i++ {
		person := GetPerson()
		chanPeople <- person
	}
	// 如果channel不被关闭的话，遍历channel 会被关闭的
	close(chanPeople)
	for value := range chanPeople {
		fmt.Println(value.Name)
	}

	fmt.Println("==============================")
	closeChannel()

}

type Person struct {
	Name    string
	Age     int
	Address string
}

func getRand() int {
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rands.Int()
}
func GetPerson() Person {

	person := Person{
		Name:    fmt.Sprintf("Name_%d", getRand()),
		Age:     getRand(),
		Address: fmt.Sprintf("Address_%d", getRand()),
	}
	return person
}

//  channel的关闭，
//   使用内置函数close 可以关闭channel,当channel关闭后，就不能再向channel写数据了，但仍然可以从channel读取数据，
func closeChannel() {

	chanints := make(chan int, 3)
	chanints <- 12
	chanints <- 13
	//关闭channel
	close(chanints)
	//此时不能再写入数据到channel中了
	fmt.Println("aasa sa")
	nn := <-chanints
	fmt.Println(nn)
}

//   channeld 遍历
//    channel支持 for-range的方式 进行遍历，请注意两个细节
//    在遍历时，如果channel没有关闭，则会出现deadlock的错误
//    在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
