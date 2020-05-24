package process

import (
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/server/utils"
	"net"
)

//处理群聊消息的逻辑

type SmsProcess struct {
}

//
func (this *SmsProcess) SendGrroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	//取出mes中的内容
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err:", err)
		return
	}
	//将数据原封不动的转发出去
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err: ", err)
		return
	}

	//遍历服务器端的onlineUsersMap 将消息转发出去
	for key, value := range userMgr.onlineUsersMap {

		if key == smsMes.UserId { //不让信息同步到自身
			continue
		}
		this.SendMesToEachOnlineUser(data, value.Conn) //转发数据

	}

}

// 转发到在线用户列表
func (this *SmsProcess) SendMesToEachOnlineUser(info []byte, conn net.Conn) {

	//创建一个transfer 实例 转发数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(info)
	if err != nil {
		fmt.Println(" 转发消息失败 err:", err)
	}
}
