package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// redis 基本使用
//     redis   远程字典服务器， redis性能非常搞，单机能达到15w QPS,通常适合做缓存，也可以做持久化
//             是完全开源免费的，高性能的key-value分布式内存数据库，基于内存运行并支持持久化的noSql数据库，是最热门的数据库之一，
func main() {

	//创建redis conn
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connect  fail err=", err)
		return
	}

	// 关闭 redis
	defer conn.Close()
	fmt.Printf("redis connect success %+v\n", conn)

	fmt.Println("使用 redis  Set写入数据")
	do, err := conn.Do("Set", "name", "xiaoma")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	fmt.Println("set name  success ", do)
	// 获取 写入的数据
	reply, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get name fail", err)
		return
	}
	fmt.Println("get name values,", reply)

	fmt.Println("=========使用hash 表结构============")
	RedisHashTest(conn)

	fmt.Println("===========   HMGet  获取所有        ==========")
	RedisHashBaichs(conn)

}

func RedisHashTest(conn redis.Conn) {

	_, err := conn.Do("HSet", "user", "name", "john")
	if err != nil {
		fmt.Println("hset fail err", err)
		return
	}
	_, err = conn.Do("HSet", "user", "age", 123)
	if err != nil {
		fmt.Println("hset fail err", err)
		return
	}
	reply, err := redis.String(conn.Do("HGet", "user", "name"))
	if err != nil {
		fmt.Println("get name fail", err)
		return
	}
	fmt.Println("get user name values==", reply)

}

//批量 set/get函数

func RedisHashBaichs(conn redis.Conn) {

	reply, err := redis.Strings(conn.Do("HMGet", "user", "name", "age"))
	if err != nil {
		fmt.Println("get name fail", err)
		return
	}
	fmt.Printf("get user name values==%v", reply)

}
