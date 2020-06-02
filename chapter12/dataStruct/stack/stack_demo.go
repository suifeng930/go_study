package main

import (
	"errors"
	"fmt"
)

//使用数组来模拟一个栈的使用

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, //当栈顶为 -1，表示栈为空
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	//显示
	stack.ShowStack()

}

type Stack struct {
	MaxTop int    //表示我们栈最大可存数个数
	Top    int    // 表示栈顶，因为栈顶规定
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {

	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full ")
	}
	this.Top++
	//放入输入
	this.arr[this.Top] = val
	return
}

func (this *Stack) ShowStack() {
	//先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for curTop := this.Top; curTop >= 0; curTop-- {
		fmt.Printf("arr[%d]=%d \n", curTop, this.arr[curTop])
	}

}
