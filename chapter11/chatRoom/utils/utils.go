package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_study/chapter11/chatRoom/common/message"
	"net"
)

// 封装 writePkg() 方法

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给客户端
	PkgLens := uint32(len(data))
	//7.2 将data的长度转换成一个 切片
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], PkgLens)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("con.Write fail err", err)
		return
	}
	// 发送data 数据本身
	lines, err := this.Conn.Write(data)
	if lines != int(PkgLens) || err != nil {
		fmt.Println(" conn.Write(data) fail ", err)
		return
	}
	return

}

// 封装一个readPkg  返回Message ,err
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("等待 buf 的输入")
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil { // 验证客户端发送过来的长度是否
		//todo  conn.Read(buf[:4]) 只有在客户端没有被关闭的清空下，才会被阻塞。
		fmt.Println("conn.read err=", err)
		//err = errors.New("read pkg  header error")
		return
	}
	fmt.Println("读到的buf=", this.Buf[:4])
	//根据读到的buf[:4]  转成一个uint32类型
	var PkgLen uint32
	PkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen 读取消息内容
	readLine, err := this.Conn.Read(this.Buf[:PkgLen])
	if uint32(readLine) != PkgLen || err != nil {
		fmt.Println("conn.Read(buf[:PkgLen]) fail err", err)
		err = errors.New("read pkg  body error")
		return
	}
	// 将buf 反序列化成 message
	err = json.Unmarshal(this.Buf[:PkgLen], &mes)
	if err != nil {
		fmt.Println(" json.Unmarshal(buf[:PkgLen], &mes) fail err", err)
		return
	}

	return
}

// 这里将这些方法关联到结构体中
type Transfer struct {
	//  分析应该存在那些字段
	Conn net.Conn   //定义连接
	Buf  [8096]byte //定义传输时的 缓冲字段
}
