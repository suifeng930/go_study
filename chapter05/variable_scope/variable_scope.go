package main

import (
	"log"
)

//变量作用域

// 函数内部声明的变量称之为局部变量，作用域仅限于函数内部
// 全局变量： 作用域在整个包中都有效，如果首字母大写，则作用域在整个程序都有效
// 如果变量时在一个代码块中，比如for/if中，那么这个变量的作用域就在这个代码块中有效

var intNum int = 50
var Name string = "hello"

func main() {

	log.Println("intNum ", intNum)
	log.Println("Name", Name)
	testVar()
	log.Println("代码块中的变量")
	for i := 1; i < 10; i++ {
		log.Println("i= ", i)
	}

}

func testVar() {
	// 函数内部声明的变量称之为局部变量，作用域仅限于函数内部
	intNum := 20
	log.Println(" intNUmber ==", intNum)
}
