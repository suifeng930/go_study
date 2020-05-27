package main

import (
	"errors"
	"fmt"
)

//使用数组实现队列
func main() {
	//1.初始化
	queue := &Queue{
		maxSize: 5,
		Array:   [5]int{},
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("===================输入add  表示添加数据到队列====================")
		fmt.Println("===================输入get  表示从队列中获取数据====================")
		fmt.Println("===================输入show  表示显示队列===================")
		fmt.Println("===================输入exit  表示退出====================")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加的队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("加入队列失败 ：", err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			i, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("get value :", i)
			}
		case "show":
			queue.ShowQurue()
		case "exit":
			return
		}
	}

}

type Queue struct {
	maxSize int
	Array   [5]int //数组==>模拟队列
	front   int    // 表示指向队列首部
	rear    int    // 表示指向队列的尾部
}

func (this *Queue) AddQueue(val int) (err error) {

	//先判断队列是否已经满了
	if this.rear == this.maxSize-1 {
		//队列已经满了 ； rear 指队列尾部已经满了
		return errors.New("queue full")
	}
	this.rear++                 // rear 指标后移
	this.Array[this.rear] = val //将数据追加到队列中
	return
}

func (this *Queue) ShowQurue() {
	//找到队首，然后遍历到队尾
	//队首不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.Array[i])
	}
	fmt.Println()
}

func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.front == this.rear { //队列为空
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.Array[this.front]
	return

}
