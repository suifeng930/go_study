package main

import "fmt"

// map    map [key] valuetype
// key 可为类型： bool   字符类型  string  指针  channel   还可能是 接口  结构体  数组
// valueType :  数字类型、string、map、struct
// 注意： slice  map  function  不能采用 “==”做判断

func main() {

	// map的声明和注意事项
	var mapString map[string]string

	// 再使用map前需要先make, make的作用就是给map分配数据内存空间，否则会报错 //panic: assignment to entry in nil map
	mapString = make(map[string]string, 10)
	mapString["n01"] = "宋江"   //panic: assignment to entry in nil map
	mapString["n01"] = "宋江01" //panic: assignment to entry in nil map
	mapString["n02"] = "武松"   //panic: assignment to entry in nil map
	mapString["n03"] = "李逵"   //panic: assignment to entry in nil map
	fmt.Println(mapString)
	// 注意点：
	//1. map在使用前一定要使用make() 初始化内存空间
	//2. map中的key 不能重复，如果重复了，会覆盖value的值
}
