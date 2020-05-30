package main

import "fmt"

//约瑟夫问题
func main() {
	//创建 boy
	first := AddBoy(5)
	//2 打印boy
	ShowBoy(first)
	//获取 boys
	PlayGame(first, 2, 3)

}

type Boy struct {
	No   int
	name string
	Next *Boy
}

//先编写一个函数，构成单向的环形链表，
//@param  num 小孩的个数
//@param  *Boy 第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环构建这个环形链表
	for i := 1; i <= num; i++ {

		boy := &Boy{
			No: i,
		}
		//
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first //构成循环

		} else {
			curBoy.Next = boy //
			curBoy = boy
			curBoy.Next = first //构成环形链表
		}
	}
	return first

}

func ShowBoy(firstBoy *Boy) {

	if firstBoy.Next == nil {
		fmt.Println("链表为空，......")
		return
	}
	//至少有一个小孩
	curBoy := firstBoy
	for {
		fmt.Printf("小孩编号=%d --->\t", curBoy.No)
		if curBoy.Next == firstBoy {
			break
		}
		curBoy = curBoy.Next
	}
}

//约瑟夫问题解决

// 分析思路
//        1.编写一个函数，PlayGame(firstBoy *Boy,startNO,countNo int)
//        1.最后使用一个算法，按照要求，在环形链表中留下最后一个人

func PlayGame(firstBoy *Boy, startNO, countNo int) {
	//1.需要定义一个辅助的指针，帮助删除小孩
	//2.空链表单独处理
	if firstBoy.Next == nil {
		fmt.Println("空链表，没有小孩....")
		return
	}
	count := GetBoyCount(firstBoy)
	if startNO > count {
		fmt.Println("startNo ,必须要小于小孩的总数。。。")
		return
	}
	tail := firstBoy
	//3. 让tail 指向环形链表的最后一个
	for {
		if tail.Next == firstBoy { //说明tail 到了最后的小孩
			break
		}
		tail = tail.Next
	}
	//4.让first 移动到startNO  [后面删除小孩，就以first为准]
	for i := 1; i <= startNO-1; i++ {
		firstBoy = firstBoy.Next //
		tail = tail.Next         //
	}
	fmt.Println()
	//5.开始数countNum下，然后就删除first 指向的小孩
	for {
		//开始数countNum-1 次
		for i := 1; i <= countNo-1; i++ {
			firstBoy = firstBoy.Next //
			tail = tail.Next         //
		}
		//删除first执行的小孩
		fmt.Printf("编号为：%d   出圈------->\n", firstBoy.No)
		firstBoy = firstBoy.Next
		tail.Next = firstBoy
		//如果只有一个小孩，就退出循环
		if tail == firstBoy {
			break
		}
	}
	fmt.Printf("最后出圈的编号为：%d   出圈------->\n", firstBoy.No)
}

func GetBoyCount(firstBoy *Boy) int {

	count := 0
	if firstBoy.Next == nil {
		fmt.Println("链表为空，......")
		return count
	}
	//至少有一个小孩
	curBoy := firstBoy
	for {
		count++
		//fmt.Printf("小孩编号=%d --->\t", curBoy.No)
		if curBoy.Next == firstBoy {
			break
		}
		curBoy = curBoy.Next
	}
	return count
}
