package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 将一个已经存在的文件 ，读取到另外一个文件中
func main() {

	filePath := "D:/Java/GOPATH/src/go_study/chapter08/testProject/main.go"
	writeFilePath := "D:/Java/GOPATH/src/go_study/chapter09/fileStream/writeDemo02/writeFile"

	readFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("read file err=", err)
		return
	}
	err = ioutil.WriteFile(writeFilePath, readFile, 0666)
	if err != nil {
		fmt.Println("write File fails ,the error is =", err)
		return
	}
}

// 判断文件是否存在
//
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}
