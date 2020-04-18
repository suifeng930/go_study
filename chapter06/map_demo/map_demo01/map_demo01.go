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

	// map切片
	// 如果切片的数据类型是map,则被称之为 slice of map,这样就能实现map的动态变化了

	//声明一个slice mqp
	var masters []map[string]string
	masters = make([]map[string]string, 2)

	if masters[0] == nil {
		masters[0] = make(map[string]string, 2)
		masters[0]["name"] = "六魔王"
		masters[0]["age"] = "500"
	}

	if masters[1] == nil {
		masters[1] = make(map[string]string, 2)
		masters[1]["name"] = "玉兔精"
		masters[1]["age"] = "500"
	}
	fmt.Println(masters)
	// 先定义一个 map
	newMasters := map[string]string{
		"name": "火云邪神",
		"age":  "120",
	}
	// 使用 append方法使用 其实就是将slice 的value  变成了 map ，做动态map追加，其实就是对slice追加

	masters = append(masters, newMasters)
	fmt.Println(masters)

}
