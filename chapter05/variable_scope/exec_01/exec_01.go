package main

import "log"

func main() {

	log.Println(name) //jackChen
	test01()
	test02() //tom
	test01()

	NAME := "Andre" // 等价于   var NAME string     NAME ="Andre"
	log.Println(NAME)

}

var name = " jackChen"

func test01() {

	log.Println(name) //jackChen
}
func test02() {
	name := "tom"
	log.Println(name) //tom
}
