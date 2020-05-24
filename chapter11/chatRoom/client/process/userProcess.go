package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_study/chapter11/chatRoom/client/utils"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/server/model"
	"net"
	"os"
)

type UserProcess struct {
}

// 用户登录方法
func (this *UserProcess) Login(userId int, passWard string) (err error) {

	//开始定义协议

	// 连接到服务器端
	conn, err := ConnServer("0.0.0.0:8889")
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
		//todo   登录成功，  1.显示登录成功的菜单(循环显示)  2.保持和服务器的通讯操作
		//  还需要再客户端启动一个协程，用于保持和服务器端的通讯，如果服务器有数据推送给客户端，则接收并显示在客户端的终端

		// 显示当前在线用户列表
		fmt.Println("当前在线用户列表如下：")
		for _, value := range loginResMes.UserIds {
			if value == userId { // 不显示当前用户在线
				continue
			}
			fmt.Printf("userId %d \n", value)
			// 完成客户端的onlineUsers 初始化
			user := &model.User{
				UserId:     value,
				UserStatus: message.UserOnline,
			}
			OnlineUsers[value] = user

		}
		go ServerProcessMes(conn) //接收服务端发送给客户端的数据
		for {                     //1.显示登录成功的菜单(循环显示)
			ShowMenu()
		}

	} else {
		fmt.Println("用户登录失败 :", loginResMes.Error)
	}

	return

}

func (this *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	conn, err := ConnServer("0.0.0.0:8889")
	// 延时关闭
	defer conn.Close()
	//2. 准备通过conn发送message结构体
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3. 创建一个loginMes
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将loginMes 序列化
	data, err := json.Marshal(registerMes)
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

	//发送数据到 server
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	//发送data
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息失败   err:", err)
	}
	mes, err = tf.ReadPkg() //register
	if err != nil {
		fmt.Println("tf.ReadPkg() fail:", err)
		return
	}
	//将mes 的data部分反序列化成为RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if registerResMes.Code == 200 {
		fmt.Println("用户注册成功，请重新登录...")
		os.Exit(0) //退出注册
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return

}

// 指定 server端ip
func ConnServer(address string) (conn net.Conn, err error) {
	//conn, err = net.Dial("tcp", "0.0.0.0:8889")
	conn, err = net.Dial("tcp", address)
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	return
}
