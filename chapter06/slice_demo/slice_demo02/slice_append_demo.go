package main

import "fmt"

// 使用append 内置函数，可以对切片进行动态追加
func main() {

	var slice01 []int = []int{100, 23, 345, 123}
	fmt.Println(slice01)
	// 通过append函数直接给slice01追加具体的元素
	slice02 := append(slice01, 123, 351, 155)

	fmt.Println("slice01", slice01)
	fmt.Println("slice02", slice02)
	slice01 = append(slice01, 123, 351, 155)
	// 直接追加到 slice01 完成slice 的扩容操作
	fmt.Println("slice01  ", slice01)
	// 切片append操作的底层原理分析：
	//1. 切片append操作的本质就是对数组扩容
	//2. go底层会创建一个新的数组newArr
	//3. 将slice原来包含的元素拷贝到新的数组NewArr
	//4. slice重新引用到newArr
	//5. 注意newArr是再底层来维护的，程序员不可见

	// 切片的拷贝操作
	var a []int = []int{1, 2, 6, 4}
	copyslice := make([]int, 10)
	// copy（para1,para2）参数的数据类型都是切片
	value := copy(copyslice, a)
	fmt.Println("copy values=", value)
	fmt.Println("copyslice=", copyslice)

	var slice04 []int
	var slice03 []int = []int{1, 2, 44, 5, 66}
	slice04 = slice03[:]
	var slice05 = slice04
	slice05[0] = 10
	fmt.Println("slice05=", slice05)
	fmt.Println("slice04=", slice04)
	fmt.Println("slice03=", slice03)

	// sendSlice
	var sendSlice []int = []int{13, 45, 23, 61, 556}
	fmt.Println("sendSlice=", sendSlice)
	testSlice(sendSlice)
	fmt.Println("sendSlice=", sendSlice)
}

func testSlice(slice []int) {
	slice[0] = 100
	fmt.Println("slice=", slice)
}
