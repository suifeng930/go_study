package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	AddArrays()

	//
	CompareArray()

	Arrays()
	PrintArrrays()
}

// 1.创建一个byte类型的26个元素的数组，分别放置 A~Z， 使用for循环访问所有元素并打印出来，提示：字符数据运算 A+ 1->B

func AddArrays() {
	var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = byte('A' + i)
	}
	for key := range myChars {
		fmt.Printf("mychar[%d]= %c\n", key, myChars[key])

	}
}

//2.  请求出一个数组的最大值，并得到对应的下标

func CompareArray() {
	var intArr [6]int
	for i := 0; i < len(intArr); i++ {
		intArr[i] = 123 + i
	}
	maxValue := intArr[0]
	maxLen := 0
	for i := 1; i < len(intArr); i++ {
		if maxValue < intArr[i] {
			maxValue = intArr[i]
			maxLen = i
		}
	}
	fmt.Printf("maxValue=%v maxLen=%v\n", maxValue, maxLen)
}

//3.  请求出一个数组的和及平均数， for range

func Arrays() {
	var intArr [6]int
	for i := 0; i < len(intArr); i++ {
		intArr[i] = 123 + i
	}
	sum := 0
	avg := 0.0
	for key := range intArr {
		sum += intArr[key]
	}
	avg = float64(sum) / float64(len(intArr))
	fmt.Printf("sum=%v avg=%v\n", sum, avg)
}

//4. 随机生成5个数，并将其反转打印

func PrintArrrays() {
	var intArr [5]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(intArr); i++ {

		intArr[i] = rand.Intn(100)
	}

	fmt.Println(intArr)
	//交换反转
	temp := 0
	for i := 0; i < len(intArr)/2; i++ {
		temp = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println(intArr)
}
