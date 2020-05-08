package main

import (
	"encoding/json"
	"fmt"
)

//json  字符串 反序列化为结构体
func main() {

	//反序列化 struct
	unMarshaltoStruct()
	unMarshaltoMap()
	unMarshaltoSlice()

}

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func unMarshaltoStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":435,\"birthday\":\"2011-12-4\",\"sal\":221336.123,\"skill\":\"睡觉\"}"
	//定义个monster实例
	var monstar Monster
	err := json.Unmarshal([]byte(str), &monstar)
	if err != nil {
		fmt.Println("json 反序列失败", err)
	}

	//输出反序列化后的设备
	fmt.Println(monstar)
	// {牛魔王 435 2011-12-4 221336.123 睡觉}
}

func unMarshaltoMap() {
	str := "{\"Address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"
	//定义个monster实例
	var monstar map[string]interface{}
	err := json.Unmarshal([]byte(str), &monstar)
	if err != nil {
		fmt.Println("json 反序列失败", err)
	}

	fmt.Println(monstar)
	//map[Address:火云洞 age:30 name:红孩儿]
}

func unMarshaltoSlice() {
	str := " [{\"address\":\"北京\",\"age\":32,\"name\":\"jack\"}," +
		"{\"address\":\"墨西哥\",\"age\":12,\"name\":\"tom\"}]"
	//定义个monster实例
	var monstar []map[string]interface{}
	err := json.Unmarshal([]byte(str), &monstar)
	if err != nil {
		fmt.Println("json 反序列失败", err)
	}

	fmt.Println(monstar)
	//[map[address:北京 age:32 name:jack] map[address:墨西哥 age:12 name:tom]]
}
