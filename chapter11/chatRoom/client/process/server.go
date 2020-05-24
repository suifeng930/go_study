package process

import (
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/client/utils"
	"go_study/chapter11/chatRoom/common/message"
	"net"
)
import "os"

//1.显示登录成功页面
//1.保持和服务器的通讯
//1.当读取服务器发送的消息后，就会显示在页面上

//显示登录成功后的页面
func ShowMenu() {
	fmt.Println("----------------恭喜xxx登录成功-----------------------")
	fmt.Println("1.显示在线列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.信息列表")
	fmt.Println("4.退出系统")
	fmt.Println("请选择(1~4):")
	var key int
	var content string

	//创建 一个smsProcess 实例，因此定义一个实例
	smsProcess := SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		OutPutOnlineUser()
	case 2:
		fmt.Println("你想对大家说什么？")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择 退出了系统....")
		os.Exit(0)
	default:
		fmt.Println("你的输入有误，请重新输入")

	}

}

// 和服务器端保持通讯
func ServerProcessMes(conn net.Conn, currentUser int) {

	//创建一个transfer ,不停的读取源服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	for {
		fmt.Printf(" 客户端  %d  正在等待读取服务器发送的消息\n", currentUser)
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() fail : ", err)
			return

		}
		// 如果读到消息， 又是下一步操作
		fmt.Println("mes=", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 有人上线了
			// 1.取出用notifyMessage
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes) err:", err)
				return
			}
			// 2.把这个用户的信息，保存在客户端的map中
			UpdateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//有人群发消息，客户端接收数据
			OutPutGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了位置的消息类型")

		}

	}

}
