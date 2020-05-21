package process

import (
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/utils"
	"net"
)

// 处理用户登录相关的handler
//           1.处理和用户相关的请求
//           2.用户登录
//           3.用户注册
//           4.用户注销
//           5.用户列表管理
type UserProcess struct {
	Conn net.Conn // tcp 连接

}

//编写处理登录请求的函数
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes。Data  并直接反序列化成 loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) fail:", err)
		return
	}
	// 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 再声明一个loginResMes
	var loginResMes message.LoginResMes

	// 2.判断用户是否合法
	if loginMes.UserId == 1000 && loginMes.UserPwd == "123456" {
		// 用户合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请注册再使用"
	}
	//3. 序列化 返回数据结构体
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) fail : ", err)
		return
	}

	//4. 将data 赋值给resMes
	resMes.Data = string(data)
	//5.对ResMes 序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail: ", err)
		return
	}

	//todo  6.将封装好的数据 返回到客户端
	// todo   抽离 writePkg()
	tf := &utils.Transfer{
		Conn: this.Conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(" writePkg(conn, data) fail ", err)
		return
	}
	return

}
