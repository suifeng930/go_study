package main

import "fmt"

func main() {

	//1.string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@qqasdas"
	//使用切片获取到qq
	sliceStr := str[6:]
	fmt.Println("sliceStr  切片string value值为：", sliceStr)
	//2. string是不可变的，也就是说不能通过 str[1]="z" 方式来修改字符串

	//3. 需要修改字符串，可以先将string--->【】byte 数组/ []rune  ---修改字符串------》重写转成string

	var arr = []byte(str)
	arr[0] = 'z'
	str = string(arr)
	fmt.Println("str=", str)

	StrCN := "Hello中国"

	arrStr := []rune(StrCN)
	arrStr[0] = '爱'
	StrCN = string(arrStr)
	fmt.Println("使用rune 转换中文字符串", StrCN)

	// 测试fbn（）方法

	fbnSlice := fbn(10)
	fmt.Println(fbnSlice)
}

//编写一个函数 fbn( n  int)  要求完成：
//1. 可以接受一个 n int
//2. 能够将斐波那契数列放到切片中
//3. 提示，斐波那契数列的数列形式  arr[0]=1 arr[1]=1 arr[2]=2 arr[3]=3 arr[4]=5 arr[5]=8
func fbn(n int) []uint64 {

	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
