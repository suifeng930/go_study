package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// copy 图片文件
func main() {
	// 将srcFile  写入到 dstFile
	srcFile := "D:\\英雄时刻\\2481827\\英雄时刻_20191020-15点39分02s.avi"
	dstFile := "D:/英雄时刻/avi.avi"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Println("copy File is Fail err: ", err)
	}

}

// 提供一个方法  接受两个文件路径 srcFileName  dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("openFile err=%v\n", err)
		return
	}
	//通过srcFile 获取reader
	reader := bufio.NewReader(srcFile)
	defer srcFile.Close()
	// 打开dstFIleName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open File err=%v", err)
		return
	}
	//通过dstFile 获取Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)

}
