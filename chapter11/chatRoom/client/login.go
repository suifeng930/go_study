package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/utils"
	"net"
)

//写一个函数， 完成登录校验
func Login(userId int, passWard string) (err error) {

	//开始定义协议
	//fmt.Printf("userId =%d userpass =%s\n", userId, passWard)
	//
	//return nil

	// 连接到服务器端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}
	// 延时关闭
	defer conn.Close()

	//2. 准备通过conn发送message结构体
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. 创建一个loginMes
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = passWard
	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println(" json.Marshal err=", err)
		return err
	}

	//5 把data 付给mess.Data字段
	mes.Data = string(data)
	//6.将mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal err=", err)
		return err
	}

	//7.data就是要发送的message
	//7.1 先把data的长度，发送给服务器
	PkgLens := uint32(len(data))
	//7.2 将data的长度转换成一个 切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], PkgLens)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("con.Write fail err", err)
		return
	}
	fmt.Printf("客户端发送的消息长度 ok.....长度 %d  内容 %v \n", len(data), string(data))
	// 8. 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(" conn.Write(data) fail ", err)
		return
	}
	// 这里还需要处理服务器端返回的消息
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	mes2, err := tf.ReadPkg()
	if err != nil {
		fmt.Println(" utils.ReadPkg(conn) fail ", err)
		return
	}
	//mes2.Data 将mes.data 反序列化成为loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes2.Data), &loginResMes)
	if err != nil {
		fmt.Println(" json.Unmarshal([]byte(mes2.Data), &loginResMes) fail ", err)
		return
	}

	if loginResMes.Code == 200 {
		fmt.Println(" 用户登录成功")
	} else {
		fmt.Println("用户登录失败 :", loginResMes.Error)
	}

	return

}

//实现用户登录
