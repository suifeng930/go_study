package main

//   海量用户即时通讯系统

//  需求分析：
//           1.用户注册
//           2.用户登录
//           3.显示在线用户列表
//           4.群聊(广播)
//           5.点对点聊天
//           6.离线留言
//  界面设计：
//
//
//

//  实现用户登录
//        1.完成客户端可以发送消息长度，服务端可以正常接收到该长度值
//        2. 完成客户端可以发送消息本身，服务器端可以正常接受到消息，并根据客户端发送的消息，判断用户的合法性，并返回响应的loginResMes
//                         1.先让客户端发送消息本身，
//                         2.服务器端接收到消息，然后反序列化成对应的消息结构体
//                         3.服务器端根据反序列化成为对应的消息，判断是否登录用户是合法，返回loginResMes
//                         4.客户端解析返回的loginResMes,显示对应界面
//                         5.这里需要对一些函数的封装
//
//
func main() {

}
