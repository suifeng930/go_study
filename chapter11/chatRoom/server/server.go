package main

import (
	"fmt"
	"go_study/chapter11/chatRoom/server/model"
	"go_study/chapter11/chatRoom/server/process"
	"net"
	"time"
)

// 负责 ：
//      1.监听Server端口
//      2.等待客户端的连接
//      3.初始化部分工作

func main() {

	// 初始化 redis 连接池
	model.InitPool("localhost:6379", 300*time.Second, 16, 0)
	//初始化 userDao
	InitUserDao()

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
		go Process(conn)

	}
}

//处理和客户端的通讯
func Process(conn net.Conn) {
	// 读取客户端发送的信息
	//  todo 这里需要延时关闭
	defer conn.Close()
	//创建一个总控开关
	processor := &process.Processor{Conn: conn}
	err := processor.Handler()
	if err != nil {
		fmt.Println(" 客户端与服务端通讯协程错误 processor.Handler() fail", err)
		return
	}

}

//完成对UserDao的初始化任务
func InitUserDao() {
	//redis pool  来自于 initRedisPool
	model.MyUserDao = model.NewUserDao(model.Pool)
}
