package main

import "fmt"

//单向环形链表
func main() {

	//  这里我们初始化一个头结点
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "白猫",
	}
	//cat2 := &CatNode{
	//	no:   2,
	//	name: "白猫",
	//}

	InsertNode(head, cat1)
	ListCircleLink(head)

	head = DeleteNode(head, 2)
	ListCircleLink(head)

}

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertNode(head *CatNode, newNode *CatNode) {
	if head.next == nil { //说明是头结点
		head.no = newNode.no
		head.name = newNode.name
		head.next = head //todo  这里指向head
		return
	}
	temp := head
	for {
		if temp.next == head { //说明找到尾部了
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = head
}

func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空的 环形链表....")
		return
	}
	for {
		fmt.Printf("cat message[no=%d name=%s] :-->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}

}

// 删除一个结点
func DeleteNode(head *CatNode, deleteId int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil { //空链表
		fmt.Println("这是一个空的环形链表，无法删除")
		return head
	}
	// 如果只有一个结点
	if temp.next == head {
		if temp.no == deleteId {

			temp.next = nil // 将头结点置为nil
		}
		return head
	}
	//将helper定位当链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//有两个以上的结点
	flag := true
	for {
		if temp.next == head { //如果到这里，说明已经比较到了最后一个元素
			break
		}
		if temp.no == deleteId { //找到结点
			if temp == head { //删除的是头结点
				head = head.next //将头结点换位
			}
			helper.next = temp.next //删除结点
			fmt.Printf("[no%d 被delete\n]", deleteId)
			flag = false
			break
		}
		temp = temp.next     //比较是不是要要删除的结点
		helper = helper.next //找到删除的结点后，删除该结点
	}
	//这里还要比较一次
	if flag { //如果flag 为true ,在for循环中没有删除元素
		if temp.no == deleteId {
			helper.next = temp.next //删除结点
			fmt.Printf("[no%d 被delete\n]", deleteId)
		} else {
			fmt.Printf("对不起，没有no%d\n", deleteId)
		}

	}
	return head

}
