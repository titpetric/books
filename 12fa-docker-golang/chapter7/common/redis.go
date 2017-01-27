package common

import "time"
import "github.com/garyburd/redigo/redis"

var (
	connectTimeout = redis.DialConnectTimeout(time.Second)
	readTimeout    = redis.DialReadTimeout(time.Second)
	writeTimeout   = redis.DialWriteTimeout(time.Second)
)

func GetRedis() (redis.Conn, error) {
	return redis.Dial("tcp", "redis:6379", connectTimeout, readTimeout, writeTimeout)
}
