package main

import (
	"bufio"
	"fmt"
	"os"
)

// 写文件的基本使用
func main() {
	filePath := "D:/GO/test.txt"
	// 创建一个新文件，写入内容为 hello ,goland
	createFile(filePath)
	// 打开一个存在的文件，将原来的内容覆盖成新的内容。  "你好，重庆"
	fmt.Println("================ 打开一个存在的文件，将原来的内容覆盖成新的内容===============")
	FlushFile(filePath)
	fmt.Println("================ 打开一个存在的文件，在文本末尾追加一段内容===============")
	AppendFile(filePath)
	fmt.Println("================ 打开一个存在的文件， 将原来内容读出显示在终端，并且追加几句话===============")
	ReadAndAppendFile(filePath)

}

// 创建一个新文件，写入一句内容
func createFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	// 及时关闭file 句柄
	defer file.Close()

	// 准备写入的文件内容
	str := "hello, goland \r\n"
	//  写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer 是带缓存的，因此在调用writeString方法时，其实内容是先写入到缓冲区，在调用flush()  将缓冲区的数据写到到文件中
	writer.Flush()

}

// 打开一个存在的文件，将原来的内容覆盖成新的内容。  "你好，重庆"
func FlushFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	str := " 你好 ，重庆 \r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer 是带缓存的，因此在调用writeString方法时，其实内容是先写入到缓冲区，在调用flush()  将缓冲区的数据写到到文件中
	writer.Flush()
}

// 打开一个文件，在文件后面追加一句新的内容，" ANC  asd "
func AppendFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	str := " 你好，北京 \r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer 是带缓存的，因此在调用writeString方法时，其实内容是先写入到缓冲区，在调用flush()  将缓冲区的数据写到到文件中
	writer.Flush()
}

//打开一个存在的文件， 将原来内容读出显示在终端，并且追加几句话
func ReadAndAppendFile(filePath string) {
	//以读写的方式打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		readStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read file err=", err)
			break
		}
		//打印在控制台
		fmt.Print(readStr)
	}
	str := " 你好，这是新写入的文件 \r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer 是带缓存的，因此在调用writeString方法时，其实内容是先写入到缓冲区，在调用flush()  将缓冲区的数据写到到文件中
	writer.Flush()
}
