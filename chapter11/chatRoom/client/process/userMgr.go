package process

import (
	"fmt"
	model2 "go_study/chapter11/chatRoom/client/model"
	"go_study/chapter11/chatRoom/common/message"
	"go_study/chapter11/chatRoom/server/model"
)

//处理服务器端的用户操作

//客户端维护的map
var OnlineUsers map[int]*model.User = make(map[int]*model.User, 10)

var CurUser model2.CurUser //在用户登录成功后，完成对curUser的初始化操作

//编写一个方法处理，服务器返回的信息
func UpdateUserStatus(mes *message.NotifyUserStatusMes) {

	user, ok := OnlineUsers[mes.UserId]
	if !ok { //原来没有用户
		user = &model.User{
			UserId: mes.UserId,
		}
	}

	user.UserStatus = mes.Status
	OnlineUsers[mes.UserId] = user

	//显示当前用户列表
	OutPutOnlineUser()

}

//显示当前在线用户
func OutPutOnlineUser() {
	//遍历一般onlineUsers
	fmt.Println("当前在线用户列表：")
	for key, _ := range OnlineUsers {
		fmt.Printf("用户id:%d\t \n", key)

	}

}
