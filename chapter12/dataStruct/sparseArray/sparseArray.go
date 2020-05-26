package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// val 结构体
type ValNode struct {
	row int
	col int
	val int
}

// 稀疏数组应用
func main() {

	//  1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	chessMap[4][3] = 1 //白子
	chessMap[2][6] = 2 //白子

	// 2. 输出看看原始的数组
	fmt.Println("1. 输出看看原始的数组")
	for _, value := range chessMap {
		for _, V := range value {
			fmt.Printf("%d\t", V)
		}
		fmt.Println()
	}

	//3. 将数组转化为稀疏数组
	fmt.Println("2. 将数组转化为稀疏数组")
	// 4.思路：
	//        1.遍历 chessMap  如果发现有个元素的值不等于0， 创建一个node结构体
	//        2. 将数据存放再结构体中
	var sparseArr []ValNode

	// 1.标准的稀疏数组应该还有一个 记录元素的二维数组的规模( 行 和列，默认值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0}
	// 2.遍历原始数组，并添加到node节点中
	sparseArr = append(sparseArr, valNode)
	for k1, v1 := range chessMap {
		for k2, v2 := range v1 {
			if v2 != 0 { // 创建一个值节点
				valNode := ValNode{
					row: k1,
					col: k2,
					val: v2}
				sparseArr = append(sparseArr, valNode) //将数据存在node节点中
			}
		}

	}

	// 4.  输出稀疏数组
	fmt.Println("3.  打印稀疏数组。。。。")
	for key, value := range sparseArr {
		fmt.Printf("%d:\t%d\t%d\t%d\t\n", key, value.row, value.col, value.val)
	}
	fmt.Println("4. 将数组存放到文件中去：")
	dstPath := "go_study/chapter12/dataStruct/sparseArray/chessMap.data"
	writeString := ""
	for _, value := range sparseArr {
		writeStr := fmt.Sprintf("%d|%d|%d|\r\n", value.row, value.col, value.val)
		writeString += writeStr
	}
	// 将数据写入到文件中go_study/chapter12/dataStruct/sparseArray
	WriteToFile(dstPath, writeString)
	fmt.Println("5. 将数组从文件中取出：")
	nodes := ReadToFile(dstPath)
	fmt.Println("取出文件中数据，并将数据存放到node中去", nodes)
	//将数据从文件中读取，并恢复原始数据
	fmt.Println("6. 将数据从文件中读取，并恢复原始数据.....")
	////   a 先创建原始数组
	var chessMap2 [11][11]int
	////   b 遍历文件的每一行
	fmt.Println("7. 遍历node ,将数据回填到二维数组中.....")
	for key, node := range nodes {
		if key == 0 {
			continue //跳过第一组数据
		}
		chessMap2[node.row][node.col] = node.val //为原始数组赋值
	}
	////    c 遍历原始数据
	fmt.Println("8 .恢复后的原始数据.....")
	for _, value := range chessMap2 {
		for _, V := range value {
			fmt.Printf("%d\t", V)
		}
		fmt.Println()
	}
}

//@dstPath  文件存放地
//@writeString  写入字符串
func WriteToFile(dstPath string, writeString string) {
	file, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(writeString)
	writer.Flush()
}

func ReadToFile(filePath string) (nodes []ValNode) {
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
			break
		}
		node := SpiltStringToNode(readStr)
		nodes = append(nodes, node)
	}
	return
}

//将读到的每一行 封装到 valNode中
func SpiltStringToNode(str string) (node ValNode) {
	split := strings.Split(str, "|")
	intRow, _ := strconv.Atoi(split[0])
	intCol, _ := strconv.Atoi(split[1])
	intVal, _ := strconv.Atoi(split[2])
	node.row = intRow
	node.col = intCol
	node.val = intVal
	return
}
