package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// tcp  client
func main() {

	conn, err := net.Dial("tcp", "192.168.1.3:8888")
	if err != nil {
		fmt.Println("tcp client err=", err)

		return
	}
	fmt.Println("tcp client  conn=", conn)
	// 1. 向server端发送数据
	reader := bufio.NewReader(os.Stdin)

	//从终端读取一行用户输入，并准备发送给服务端
	fmt.Println("请输入你要发送的数据，输入exit/EXIT 退出输入")
	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string fail err=", err)
			return
		}
		if line == "exit\n" || line == "EXIT\n" {
			fmt.Println("你输入的是exit or EXIT，程序将主动退出")
			break
		}

		//将数据发送给服务端
		num, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Wrrite err=", err)
		}
		if num > 0 {
			fmt.Printf("conn.Write()  success... send  %d 字节数据 \n", num)
		}
		time.Sleep(1 * time.Second)
	}

}
