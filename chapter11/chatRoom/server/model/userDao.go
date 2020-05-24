package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个UserDao struct  完成对User 结构体的各种操作

//定义个全局UserDao实例
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

// todo  使用工厂模式哦，创建一个UserDao实例
func NewUserDao(redisPool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: redisPool,
	}
	return
}

//根据用户id  返回一个user实例
func (this *UserDao) GetUserById(conn redis.Conn, userId int) (user *User, err error) {

	//  通过给定的id , 去redis 查询这用户
	reply, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		// 报错
		if err == redis.ErrNil {
			err = ERROR_USER__NOTEXIST //用户不存在
		}
		return
	}
	// 将数据反序列化成user
	err = json.Unmarshal([]byte(reply), &user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(reply), &user) fail :", err)
		return
	}
	return
}

// 完成登录的校验
//  1.login  完成对用户的验证
//  2.如果用户的id和pws 都正确，则返回一个user 实例
//  3.如果用户的id和pws 有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPws string) (user *User, err error) {

	//先获取 userDao 的连接池
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	//判断用户密码是否正确
	if user.UserPwd != userPws {
		err = ERROR_USER__Pwd
		return
	}
	return

}

//用户注册功能
func (this *UserDao) Register(user *User) (err error) {
	//先从userDao 的连接池取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER__EXISTS //用户存在
		return
	}
	// 用户不存在，支持用户注册操作
	data, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(" json.Marshal(&user) err:", err)
		return
	}
	// 完成入库操作
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误，err=", err)
		return
	}
	return

}
