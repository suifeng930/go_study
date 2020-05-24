package process

import (
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/server/model"
	"go_study/chapter11/chatRoom/server/utils"
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
	//  去redis 中查询数据 验证数据的合法性
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		// 存在错误
		//不合法
		// todo    我们先测试成功，然后再根据返回具体错误信息
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请注册再使用"
	} else {
		// 用户合法
		loginResMes.Code = 200

	}
	fmt.Println("用户登录成功 ：", user)
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
