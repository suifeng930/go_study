package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"io"
	"net"
)

func main() {

	fmt.Println("服务器在8880端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	//这里需要延时关闭  todo
	defer listen.Close()

	//一旦监听成功，就等待客户端链接服务器
	for {
		fmt.Println("等待客户端来连接服务器。。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(" listen.Accept() err =", err)

		}

		//一旦连接成功，则启动一个协程和客户端 保存通讯。。。
		go process(conn)

	}
}

//处理和客户端的通讯
func process(conn net.Conn) {
	// 读取客户端发送的信息

	//这里需要延时关闭  todo
	defer conn.Close()
	for {
		mes, err := ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭，服务器端也退出。。。。")
				return
			} else {
				fmt.Println("ReadPkg(conn) fail", err)
				return

			}
		}
		fmt.Println("mes=", mes)
	}

}

// 封装一个readPkg  返回Message ,err
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("等待 buf 的输入")
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil { // 验证客户端发送过来的长度是否
		//todo  conn.Read(buf[:4]) 只有在客户端没有被关闭的清空下，才会被阻塞。
		fmt.Println("conn.read err=", err)
		//err = errors.New("read pkg  header error")
		return
	}
	fmt.Println("读到的buf=", buf[:4])
	//根据读到的buf[:4]  转成一个uint32类型
	var PkgLen uint32
	PkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen 读取消息内容
	readLine, err := conn.Read(buf[:PkgLen])
	if uint32(readLine) != PkgLen || err != nil {
		fmt.Println("conn.Read(buf[:PkgLen]) fail err", err)
		err = errors.New("read pkg  body error")
		return
	}
	// 将buf 反序列化成 message
	err = json.Unmarshal(buf[:PkgLen], &mes)
	if err != nil {
		fmt.Println(" json.Unmarshal(buf[:PkgLen], &mes) fail err", err)
		return
	}

	return
}
