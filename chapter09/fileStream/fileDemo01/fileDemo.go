package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
)
import "os"

//文件的基本操作
func main() {

	//打开文件的方法

	fileName := "D:/Java/GOPATH/src/go_study/chapter08/testProject/main.go"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)

	}
	//输入文件内容
	fmt.Printf("file=%v\n", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)

	}
	fmt.Println("===============================================")
	bufferRead(fileName)
	fmt.Println("==================ioutil.ReadFile()========================")
	ReadFileByOnce(fileName)

}

//读取文件的内容，显示到终端(带缓冲区的方法)  os.Open()  file.Close()  buffer.newReader()

func bufferRead(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏
	// 创建一个reader 带缓冲区的独写
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行符就就结束当前行
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str) //打印文件类容
	}
	fmt.Println("文件读取结束。。。")
}

//读取文件的内容并显示在终端(使用ioutiluoco一次将整个文件读入到内存中)，这种方式适用于文件不大的情况，相关方法和函数(ioutil.ReadFile)
//适用于文件不太大的情况
func ReadFileByOnce(fileName string) {
	readFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file err=%v \n", err)
	}
	//把文件拂去到的内容显示在终端
	fmt.Printf("%v\n", readFile)

}
