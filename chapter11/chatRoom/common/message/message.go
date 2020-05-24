package message

import "go_study/chapter11/chatRoom/server/model"

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` //消息的类型
}

//
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //密码
	UserName string `json:"userName"` //用户名

}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码 500 表示用户未注册  200 表示登录成功
	Error string `json:"error"` //返回错误信息

}

//用户注册结构体
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 用户已存在  200 表示登录成功
	Error string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User model.User `json:"user"` //注册用户结构体
}
