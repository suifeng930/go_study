package main

import (
	"bufio"
	"fmt"
	"io"
)
import "os"

//   统计文件中英文， 数字， 空格 和其他字符的数量

var reader *bufio.Reader

func main() {

	// 打开一个包含 中英文 数字 空格的文件

	// 1.创建一个reader
	// 2, 每读取一行，就去统计该行有多少个 英文 数字 空格 和其他字符
	// 3. 然后将结果保存到一个结构体中

	fileName := "D:/英雄时刻/countFIle.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	var count CharCOunt
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历
		for _, value := range readString {
			switch { //当成 if / else
			case value >= 'a' && value <= 'z':
				fallthrough //穿透
			case value >= 'A' && value <= 'Z':
				count.ChCount++
			case value == ' ' || value == '\t':
				count.SpaceCount++
			case value >= '0' || value <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的英文个数为%d,数字的个数为：%d,空格的数字为%d,其他字符为：%d", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}

//定义一个结构体，用于保存统计结果
type CharCOunt struct {
	ChCount    int // 记录英文的个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int //记录其他字符的个数

}
