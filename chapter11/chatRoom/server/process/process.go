package process

import (
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/server/utils"
	"io"
	"net"
)

// server端总的处理器
//              1.根据客户端的请求，调用对应的处理器，完成响应的任务操作
type Processor struct {
	Conn net.Conn //tcp 连接
}

// 编写一个ServerProcess()  根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) ServerProcessMessage(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		userProcess := &UserProcess{
			Conn: this.Conn,
		}
		err := userProcess.ServerProcessLogin(mes)
		return err
	case message.RegisterResMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存咋，无法处理...")
	}
	return
}

func (this *Processor) Handler() (err error) {
	for {
		//创建一个Transfer
		tf := &utils.Transfer{
			Conn: this.Conn,
			Buf:  [8096]byte{},
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭，服务器端也退出。。。。")
				return err
			} else {
				fmt.Println("ReadPkg(conn) fail", err)
				return err

			}
		}
		//fmt.Println("mes=", mes)
		err = this.ServerProcessMessage(&mes)
		if err != nil {
			return err
		}

	}
}
