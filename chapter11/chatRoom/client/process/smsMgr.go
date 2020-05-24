package process

import (
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
)

func OutPutGroupMes(mes *message.Message) {
	//这个数据一定是来着于mes

	//1.反序列化数据
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println(" json.Unmarshal([]byte(mes.Data), &smsMes) err:", err)
		return
	}

	//显示消息信息
	fmt.Printf("用户id:%d\t 对大家说：\t%s \n\n", smsMes.UserId, smsMes.Content)

}
