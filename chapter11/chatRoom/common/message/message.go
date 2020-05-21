package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
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

type RegisterResMes struct {

	// 注册
}
