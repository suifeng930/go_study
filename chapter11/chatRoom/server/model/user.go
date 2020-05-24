package model

//定义一个用户结构体，用于存入DB
type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //记录用户在线状态
}
