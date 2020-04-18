package main

import (
	"fmt"
	"sort"
)

func main() {

	//map 的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 98
	map1[4] = 56
	map1[8] = 13

	fmt.Println(map1)
	// 如果按照map 的key的顺序排列输出
	//1. 先将map.key  放入到切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照key值来输出map的值

	var keys []int
	for key, _ := range map1 {
		keys = append(keys, key)
	}
	//排序 key 的值
	sort.Ints(keys)
	fmt.Println("map 的key 排序：", keys)
	// 遍历输出map1
	for _, value := range keys {
		fmt.Printf("map1[%v]=%v \n", value, map1[value])
	}

	// map的使用细节

	// 1. map是引用类型，遵守引用类型传递的机制，在一个函数接受map,修改后，会直接修改原来的map
	modify(map1)
	fmt.Println("modify map :", map1)
	//2. map的容量达到后，在想用map,追加元素，会自动扩容，不会发生panic，也就是说map可以动态扩容
	//3. map的value也经常使用struct类型，更适合管理一些复杂的数据

}

func modify(map1 map[int]int) {
	fmt.Println("before map: ", map1)
	map1[10] = 123
	fmt.Println("after map: ", map1)

}
