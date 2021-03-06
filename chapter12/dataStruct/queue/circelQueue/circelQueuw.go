package main

import (
	"errors"
	"fmt"
)

//环形队列
func main() {

	queue := &CircleQueue{
		maxSize: 5,
		Array:   [5]int{},
		head:    0,
		tail:    0,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Println("加入队列失败 ：", err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			i, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("get value :", i)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			return
		}
	}
}

type CircleQueue struct {
	maxSize int
	Array   [5]int //数组==>模拟队列
	head    int    // 表示指向队列首部
	tail    int    // 表示指向队列的尾部
}

//进队列
func (this *CircleQueue) Push(val int) (err error) {

	// 先判断队列是否满了
	if this.IsFull() {
		return errors.New(" queue full")
	}
	this.Array[this.tail] = val                //把值给尾部
	this.tail = (this.tail + 1) % this.maxSize // tail 在队尾，但是不包含最后的元素
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	// 判断队列是否为空值
	if this.IsEmpty() {
		return -1, errors.New("queue empty")
	}
	//取出 head是指向队首，并且含队首元素
	val = this.Array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return

}

//显示队列

func (this *CircleQueue) ListQueue() {
	// 取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println(" 队列为空")
	}
	//设计一个辅助标量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.Array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()

}

//判断环形队列为满了
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head

}

// 判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.head == this.tail

}

//取出环形队列中到底有多少元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize

}
