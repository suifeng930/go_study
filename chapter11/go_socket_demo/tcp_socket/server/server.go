package main

import (
	"fmt"
	"io"
	"net"
)

// server 端代码

func main() {

	fmt.Println("服务器开始监听........")
	// 1. tcp 指定监听协议
	// 2. tcp 指定tcp监听的端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen fails -", listen)
		return
	}
	defer listen.Close() //延时关闭listen
	for {

		//等待客户端的连接
		fmt.Println("等待客户端来连接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept()  err=", err)

		} else {
			fmt.Printf("Accept() success conn=%v   client Ip=%v\n", conn, conn.RemoteAddr())
		}
		//这里准备其中一个协程，为客户端服务
		go Process(conn)

	}
	fmt.Println("listen success -", listen)
}

func Process(conn net.Conn) {

	//循环接受客户端 发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		// todo 等待客户端通过 conn 发送信息，如果客户端没有发送数据，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%v发送的信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn 读取
		if err != nil {
			fmt.Println("服务器端的read()  err=", err)
			return
		}
		if err == io.EOF {
			fmt.Println("客户端退出连接，err", err)
			return //退出函数
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}

}
