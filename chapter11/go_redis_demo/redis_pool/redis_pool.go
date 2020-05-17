package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// redis 连接池
//  1.事先初始化一定数量的连接，放入到链接池中
//  1.当go需要操作redis时，直接从redis链接池取出连接即可，
//  1.这样可以节省临时获取redis连接的时间，从而提高效率
func main() {

	//先从 pool中取出一个连接
	redisInit()
	conn := pool.Get()
	defer pool.Close() //一旦关闭连接池，就不能从连接池中再取出连接
	_, err := conn.Do("Set", "name", "tom cat")
	if err != nil {
		fmt.Println("conn.do err", err)
		return
	}
	//取出
	do, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get name fail ", err)
		return
	}
	fmt.Println("get name  ===", do)

}

var pool *redis.Pool

func redisInit() {

	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) { // 初始化连接的代码，指定连接那个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         8,   //最大空闲链接数
		MaxActive:       0,   // 表示和数据库连接的最大连接数， 0表示没有限制
		IdleTimeout:     100, //最大空闲实践，
		Wait:            false,
		MaxConnLifetime: 0,
	}
	//get := pool.Get() //从连接池中取出一个连接

}
