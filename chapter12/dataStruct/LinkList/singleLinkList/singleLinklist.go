package main

import "fmt"

//单链表实现
func main() {

	head := &HeroNode{} //创建一个头节点

	//2.创建一个新的HeroNode节点
	hero01 := &HeroNode{
		no:       1,
		name:     "松江",
		nickName: "及时雨",
	}
	hero02 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickName: "玉麒麟",
	}
	hero03 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickName: "豹子头",
	}
	hero04 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickName: "智多星",
	}
	//3.添加到链表的尾部
	//InsertHeroNode(head, hero01)
	//InsertHeroNode(head, hero02)
	//InsertHeroNode(head, hero03)
	InsertHeroNode2(head, hero03)
	InsertHeroNode2(head, hero01)
	InsertHeroNode2(head, hero02)
	InsertHeroNode2(head, hero04)
	//4. 显示链表
	ListHeroNode(head)
}

type HeroNode struct {
	no       int //排序
	name     string
	nickName string
	next     *HeroNode //指向下一个结点
}

// 1.在链表尾部追加链表
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1.先找到改链表的最后一个节点
	//2.创建一个辅助节点{跑龙套}
	temp := head
	for {
		if temp.next == nil { //找到头结点
			break
		}
		//让temp不断的指向下一个结点
		temp = temp.next
	}
	//3. 将newHeroNode加入到链表的最后
	temp.next = newHeroNode

}

// 2.根据no的编码从小到大插入...todo  实用性比较强
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//1.找到适当的结点
	//2.创建一个辅助节点{跑龙套}
	temp := head

	flag := true
	// 插入的结点no ,和temp的下一个结点的no 比较
	for flag {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode 就应该加入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明链表中已经存在这个no,就不允许插入。
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经存在 no,", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

// 显示链表的手游节点信息
func ListHeroNode(head *HeroNode) {
	// 1.创建一个辅助节点[跑龙套]
	temp := head
	//2. 先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println(" 这个链表是空链表")
		return
	}
	//3.遍历整个链表
	for {
		//打印下一个结点
		fmt.Printf("[%d,%s,%s]---->", temp.next.no, temp.next.name, temp.next.nickName)
		//判断是否在链表尾部
		temp = temp.next
		if temp.next == nil {
			break
		}
	}

}
