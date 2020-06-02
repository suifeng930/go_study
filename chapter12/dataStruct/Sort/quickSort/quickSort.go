package main

import "fmt"

//快速排序
func main() {

	arr := [6]int{-9, 78, 0, 23, -567, 70}
	fmt.Println(arr)
	QuickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)

}

//快速排序
//@param  left   数组左下标的
//@param  right   数组右下标的
func QuickSort(arr *[6]int, left, right int) {

	l := left
	r := right
	temp := 0
	pivot := arr[(left+right)/2]
	for l < r {
		for arr[l] < pivot { //从pivot 做边找大于等于 pivot的值
			l++
		}
		for arr[r] > pivot { //从pivot 右边找小于等于 pivot的值
			r--
		}
		if l >= r { //已经完成分割
			break
		}
		temp = arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	//如果l==r,再移动一下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(arr, left, r)
	}
	//向右递归
	if right > l {
		QuickSort(arr, l, right)
	}
}
