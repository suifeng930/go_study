package main

import (
	"encoding/json"
	"fmt"
)

//json 序列化

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func main() {

	// 演示 将结构体 map   切片 进行序列化
	testStruct()
	testMap()
	testSlice()
}

func testStruct() {

	// 演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      435,
		Birthday: "2011-12-4",
		Sal:      221336.123,
		Skill:    "睡觉",
	}
	//将 monster序列化
	bytes, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json maeshal fail err =", err)
		return
	}

	// 输出序列化后的结果
	fmt.Println(string(bytes))
}

// 将map进行序列化
func testMap() {
	//定义一个map
	var strMap map[string]interface{}
	strMap = make(map[string]interface{})
	strMap["name"] = "红孩儿"
	strMap["age"] = 30
	strMap["Address"] = "火云洞"
	data, err := json.Marshal(strMap)
	if err != nil {
		fmt.Println("json 序列化失败:", err)
		return
	}
	fmt.Println(string(data))

}

// 对切片进行序列化

func testSlice() {
	var slice []map[string]interface{}
	var maps map[string]interface{}
	maps = make(map[string]interface{})
	maps["name"] = "jack"
	maps["age"] = 32
	maps["address"] = "北京"
	slice = append(slice, maps)
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "tom"
	m["age"] = 12
	m["address"] = "墨西哥"
	slice = append(slice, m)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败 ", err)
		return
	}
	fmt.Println(string(data))

}
