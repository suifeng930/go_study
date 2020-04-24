package main

import "fmt"

// 继承的抽象

// 继承 提高了代码的复用性
// 代码的扩展性和可维护性提升了

//   继承的特性
//  1.结构体可以使用嵌套匿名结构体所有的字段和方法，即首字母大写或者小写的字段、方法、都可以使用
//  2.匿名结构体字段访问可以简化
//  3.当结构体和匿名结构体有相同的字段或方法时，编译器采用就近原则访问，如果希望访问匿名结构体达到字段和方法，可以通过匿名结构体名来区分
//  4.如果一个struct 嵌套了一个有名结构体，这种模式称为组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须要带上结构体的名称
//  5.嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值

func main() {

	pupil := &Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Age = 12
	pupil.Score = 143
	pupil.testing()
	pupil.SetScore(86)
	pupil.ShowInfo()

	fmt.Println("继承的特性")
	b := B{}
	b.A.Name = "Hey"
	b.A.age = 15
	b.A.hello()
	b.A.SayOk()
	fmt.Println("事实证明结构体可以使用匿名结构体的所有方法和字段")
	b.Name = "B.Name>Tom"
	b.age = 34
	fmt.Println("子类可以直接调用父类的字段和方法")

	// 如果一个struct 嵌套了一个有名结构体，这种模式称为组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须要带上结构体的名称
	fmt.Println("如果一个struct 嵌套了一个有名结构体，这种模式称为组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须要带上结构体的名称")
	var d D
	d.a.age = 23
	d.a.hello()

}

//  抽离出学生基础类
type Student struct {
	Name  string
	Age   int
	Score int
}

type Pupil struct {
	Student //嵌入匿名结构体
}
type Graduate struct {
	Student
}

// 抽离公共方法
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v  年龄=%v 成绩=%v\n ", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

//保留特有的方法
func (p *Pupil) testing() {
	fmt.Println(" 小学生正在考试。。。。。")

}
func (p *Graduate) testing() {
	fmt.Println(" 小学生正在考试。。。。。")

}

//继承细节讨论

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("a SayOk  ", a.Name)

}
func (a *A) hello() {
	fmt.Println("A Hello", a.Name)

}

//   B 是子类， A 是父类
type B struct {
	A
}

//  4.如果一个struct 嵌套了一个有名结构体，这种模式称为组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须要带上结构体的名称
type D struct {
	a A
}
