package main

import (
	"fmt"
	"go_study/chapter07/factory/model"
)

func main() {

	// 初始化一个student实例
	stu := model.NewStudent("Tome.", 78.8)

	fmt.Println("Stu :", stu) //Stu : &{Tome. 78.8}
	//fmt.Printf("stu.Name=%v  stu.score=%v \n", stu.Name, stu.Score)
	fmt.Printf("stu.Name=%v  stu.score=%v \n", stu.Name, stu.GetScore())

}
