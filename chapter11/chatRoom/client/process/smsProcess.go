package process

import (
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/client/utils"
	"go_study/chapter11/chatRoom/common/message"
)

//完成用户群发消息的操作

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 1.创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2,创建一个SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content //内容发送
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err:", err)
		return
	}
	mes.Data = string(data)
	//4.将mes  序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal(mes) err:", err)
		return
	}
	//5将mes  发送到 服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	//6 发送数据
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(" err = tf.WritePkg(data) err:", err)
		return
	}
	return
}
