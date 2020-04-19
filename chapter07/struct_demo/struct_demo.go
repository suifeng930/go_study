package main

import "fmt"

func main() {

	//1. golang 支持面向对象编程(oop),但和传统面向对象编程有区别，所有说golang支持面向对象编程特性
	//2. golang 中结构体和其他编程语言中的class 类似
	//3. golang面向对象编程非常简洁，去到了传统oop语言的继承、方法重载、构造函数、和析构函数、以及隐藏this指针等
	//4. golang仍然有面向对象编程的继承，封装，和多态的特性，只是实现的方式和其它oop语言不一样，比如继承，采用匿名字段来实现
	//5. golang面向对象很优雅，oop本身就是语言类型系统的一部分，通过接口关联interface,耦合性地，也非常灵活，

	// 创建结构体变量和访问结构体字段的形式

	//1. 类型推导

	stu01 := Student{
		Name: "xiaoma",
		Age:  15,
	}
	fmt.Println("stu01=", stu01)
	//2. 使用new
	var stu02 *Student = new(Student)
	(*stu02).Name = "小红"
	(*stu02).Age = 14
	// 另一种使用方式   stu02.Name="mary"
	stu02.Name = "mary"
	stu02.Age = 25
	fmt.Println("stu02:", *stu02)

	//3. 使用 var stu03 *Student = &Student{}
	var stu03 *Student = &Student{
		Name: "Tom",
		Age:  34,
	}
	fmt.Println("stu03:", *stu03)
}

type Student struct {
	Name string
	Age  int
}

// 如果结构体的字段类型是： 指针，slice 和map  的零值都是nil. 即还没有分配空间
// 如果需要使用这样的字段，需要先使用make,才能使用
type Person struct {
	Name  string
	Age   int
	Ptr   *int              //指针
	Slice []int             //切片
	Map01 map[string]string //map
}

//struct类型的内存分配机制
//
