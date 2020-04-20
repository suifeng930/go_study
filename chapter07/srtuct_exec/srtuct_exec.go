package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var p1 Person
	p1.Name = "小米"
	p1.Age = 14
	var p2 *Person = &p1 //创建p2 指向 p1的地址
	fmt.Println("p2.Name :", (*p2).Name)
	fmt.Println("p2.age :", p2.Age)
	// p2创建的是一个Person指针，而p2指向的是p1，因此修改p2的值也将需改p1的值
	p2.Name = "Tom"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name)
	fmt.Printf("p2.address=%p p2.value=%p p1.address=%p \n", &p2, p2, &p1)

	left := Point{
		x: 12,
		y: 5,
	}
	right := Point{
		x: 13,
		y: 15,
	}
	r1 := Rect{
		leftUp:    left,
		rightDown: right,
	}
	// r1 有四个int,在内存中是连续分布的
	fmt.Println("打印地址值：")
	fmt.Printf("r1.leftUp.x =%p,r1.leftUp.y=%p,r2.right.x=%p,r2.right.y=%p \n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// 创建一个monster 变量
	monstar := Monster{
		Name:  "六魔王",
		Age:   456,
		Skill: "牛角功",
	}
	fmt.Println("monstar ", monstar)
	//将变量序列化成 json格式
	bytes, e := json.Marshal(monstar)
	if e != nil {
		fmt.Println("序列化数据错误")
	}
	fmt.Println("bytes:", string(bytes))

}

type Person struct {
	Name string
	Age  int
}

//1. 结构体的所有字段在内存中是连续分布的， 好处方便，数据查找

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

//2.结构体是用户自定义的类型， 结构体进行type重新定义(相当于取别名)，golang认为是新的数据类型，但是互相转换必须要强转

//3. struct的每一个字段，可以写一个tag,该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}
