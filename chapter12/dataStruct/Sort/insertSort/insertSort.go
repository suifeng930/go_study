package main

import "fmt"

//插入排序
func main() {

	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println(arr)
	InsertSort(&arr)
	fmt.Println(arr)

}

func InsertSort(arr *[5]int) {

	//先完成第一个，即给第二个元素找到合适的位置并插入
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1 //插入元素下标
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--                         //指针向前走
		}
		//插入数据
		if insertIndex+1 != i {

			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后的结果：%+v \n", i, *arr)
	}

}
