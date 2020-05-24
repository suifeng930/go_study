package model

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var Pool *redis.Pool

//初始化 redis 连接池
func InitPool(address string, idleTime time.Duration, maxIdle, maxActive int) {
	Pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
		TestOnBorrow:    nil,
		MaxIdle:         maxIdle,   // 最大空闲连接数
		MaxActive:       maxActive, ///表示和数据库的最大连接数，0 表示没有限制
		IdleTimeout:     idleTime,  //超时时间
		Wait:            false,
		MaxConnLifetime: 0,
	}

}
