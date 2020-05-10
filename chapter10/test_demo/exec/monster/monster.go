package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	// 先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("json marshal fail :", err)
		return false
	}
	// 再保存到文件
	filePath := "../monster.json"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file fail ", err)
		return false
	}
	return true

}

func (this *Monster) ReStore() bool {
	// 先读取文件
	filePath := "../monster.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("write file fail ", err)
		return false
	}

	// 再封装到结构体中
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("json unmarshal fail :", err)
		return false
	}
	return true
}
