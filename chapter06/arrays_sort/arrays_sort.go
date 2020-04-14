package main

import "fmt"

// 排序算法   内部排序   外部排序

// 内部排序： 将需要处理的所有数据都加载到内部存储器中进行排序(交换式排序、选择排序、插入排序)

// 外部排序： 数据量过大，无法全部加载到内存中，需要借助外部存储进行排序(合并排序、直接合并排序)

// 交换式排序： 属于内部排序发，是运用数据值并比较后，以判断规则对数据位置进行交换，以达到排序的目的
//            冒泡排序： （bubble sort）
//            快速排序： （quick sort）
// 冒泡排序：思想-->通过对待排序序列从后向前(从下标比较大的元素开始)，依次比较相邻元素的排序码66，若发现逆序则交换，使排序码较小达到元素逐渐从后部移向前部（
//          从下标较大的单元移向下标较小的单元），就像水底的气泡一样逐渐向上冒泡。因为排序的过程中，各元素不断接近自己的位置，如果一趟比较下来没有进行过交换，
//          就说明序列有序，因此要咋排序过程中设置一个标记flag判断元素是否进行过交换。从而减少不必要的比较。

func main() {

	var intArrays [5]int = [5]int{24, 58, 76, 54, 23}
	BubbleSort(&intArrays)

	var findArrays [5]int = [5]int{1, 34, 666, 756, 849}
	BinaryFind(&findArrays, 0, len(findArrays)-1, 756)

}

//冒泡demo  将下面的数组24，59，76，12，74

func BubbleSort(arr *[5]int) {
	fmt.Println("before sort", *arr)
	temp := 0
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
		fmt.Printf("the %d sort :%v\n ", j, *arr)
	}
	fmt.Println("after sort: ", *arr)
}

// 查找    顺序查找 、  二分查找(要求数组是有序的)

//  有一个数列， 白眉鹰王 、 金毛狮王、紫衫龙王、青翼福王
//  猜数与游戏： 从键盘中任意输入一个字符，判断数列中是否包含测名称【顺序查找】

// 请对一个有序数组进行二分查找   [1,34,666,756,849]
//  先找到中间下标  middle=(left+right)/2, 然后让中间下标的值，和findValue进行比较
//  如果 arr[middle]>findValue,就应该向 leftIndex--(middle-1)
//  如果 arr[middle]>findValue,就应该向 leftIndex--(middle-1)

func BinaryFind(arr *[5]int, leftIndex, rightIndex, findValue int) {
	fmt.Println("before sort", *arr)

	if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findValue {
		BinaryFind(arr, leftIndex, middle-1, findValue)
	} else if (*arr)[middle] < findValue {
		BinaryFind(arr, middle+1, rightIndex, findValue)
	} else {
		fmt.Printf("find it , %v\n", middle)
	}

}
