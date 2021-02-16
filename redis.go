package redis

import (
	_redis "github.com/gomodule/redigo/redis"
)

type Config struct {
	Address   string // 数据库连接字符串
	Auth      string // 最大打开的连接数
	MaxIdle   int    // 最大空闲的连接数
	MaxActive int    // 连接最大的生命时间
}

var pool *_redis.Pool

func Init(conf Config) {
	pool = &_redis.Pool{
		MaxIdle:   conf.MaxIdle,
		MaxActive: conf.MaxActive,
		Dial: func() (_redis.Conn, error) {
			c, err := _redis.Dial("tcp", conf.Address)
			if err != nil {
				return nil, err
			}
			if conf.Auth != "" {
				_, err = c.Do("AUTH", conf.Auth)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}

func GetConn() _redis.Conn {
	return pool.Get()
}
