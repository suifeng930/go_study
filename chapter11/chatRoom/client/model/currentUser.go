package model

import (
	"go_study/chapter11/chatRoom/server/model"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser   我们将其定义为一个全局变量
type CurUser struct {
	Conn net.Conn
	model.User
}
