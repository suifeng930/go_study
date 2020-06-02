package main

import "fmt"

//选择排序
func main() {

	arr := [5]int{10, 34, 19, 100, 86}
	fmt.Println(arr)
	fmt.Println()
	selectSort(&arr)

}

func selectSort(arr *[5]int) {

	//先检索出第一个最大值， arr[0]
	//先假设第一个元素就是最大值
	for j := 0; j < len(arr)-1; j++ {

		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次循环：%+v\n", j, *arr)
	}
}
