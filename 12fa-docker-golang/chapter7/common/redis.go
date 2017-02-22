package common

import "time"
import "github.com/garyburd/redigo/redis"

type Redis struct {
	conn                                      redis.Conn
	address                                   string
	connectTimeout, readTimeout, writeTimeout time.Duration
}

type RedisOption func(*Redis)

func RedisAddress(address string) RedisOption {
	return func(do *Redis) {
		do.address = address
	}
}

func RedisConnectTimeout(timeout time.Duration) RedisOption {
	return func(do *Redis) {
		do.connectTimeout = timeout
	}
}

func RedisReadTimeout(timeout time.Duration) RedisOption {
	return func(do *Redis) {
		do.readTimeout = timeout
	}
}
func RedisWriteTimeout(timeout time.Duration) RedisOption {
	return func(do *Redis) {
		do.writeTimeout = timeout
	}
}

func NewRedis(options ...RedisOption) *Redis {
	redis := &Redis{
		address:        "redis:6379",
		connectTimeout: time.Second,
		readTimeout:    time.Second,
		writeTimeout:   time.Second,
	}
	for _, option := range options {
		option(redis)
	}
	return redis
}

var (
	redisConnections = make(map[string]*Redis, 0)
)

func (t *Redis) Save(names ...string) {
	connection := "default"
	for _, name := range names {
		connection = name
	}
	redisConnections[connection] = t
}

func GetRedis(names ...string) (redis.Conn, error) {
	var err error
	connection := "default"
	for _, name := range names {
		connection = name
	}
	r, ok := redisConnections[connection]
	if !ok {
		r = NewRedis()
		r.Save(connection)
	}
	if r.conn == nil {
		connectTimeout := redis.DialConnectTimeout(r.connectTimeout)
		readTimeout := redis.DialReadTimeout(r.readTimeout)
		writeTimeout := redis.DialWriteTimeout(r.writeTimeout)
		r.conn, err = redis.Dial("tcp", r.address, connectTimeout, readTimeout, writeTimeout)
	}
	return r.conn, err
}
