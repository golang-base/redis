package token

import (
	"github.com/go-basic/uuid"
	"github.com/golang-base/redis"
	_redis "github.com/gomodule/redigo/redis"
	"strings"
)

var Expire = 60 * 60 * 24 // 失效时长，秒，默认24小时
var Prefix = "token:"     // 前缀

func Set(value string) (token string, err error) {
	conn := redis.GetConn()
	defer conn.Close()

	token = genToken()
	_, err = conn.Do("SET", Prefix+token, value, "EX", Expire)
	return
}

func Get(token string) (value string, err error) {
	conn := redis.GetConn()
	defer conn.Close()

	conn.Do("GET", Prefix+token)
	value, err = _redis.String(conn.Do("GET", Prefix+token))
	return
}

func Del(token string) (err error) {
	conn := redis.GetConn()
	defer conn.Close()

	_, err = conn.Do("DEL", Prefix+token)
	return
}

func genToken() string {
	return strings.ReplaceAll(uuid.New(), "-", "")
}
