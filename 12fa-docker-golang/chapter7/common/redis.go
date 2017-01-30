package common

import "time"
import "github.com/garyburd/redigo/redis"

type Redis struct {
	conn                                      redis.Conn
	protocol, address                         string
	connectTimeout, readTimeout, writeTimeout time.Duration
}

type RedisOption struct {
	f func(*Redis)
}

func RedisProtocol(protocol string) RedisOption {
	return RedisOption{func(do *Redis) {
		do.protocol = protocol
	}}
}

func RedisAddress(address string) RedisOption {
	return RedisOption{func(do *Redis) {
		do.address = address
	}}
}

func RedisConnectTimeout(timeout time.Duration) RedisOption {
	return RedisOption{func(do *Redis) {
		do.connectTimeout = timeout
	}}
}

func RedisReadTimeout(timeout time.Duration) RedisOption {
	return RedisOption{func(do *Redis) {
		do.readTimeout = timeout
	}}
}
func RedisWriteTimeout(timeout time.Duration) RedisOption {
	return RedisOption{func(do *Redis) {
		do.writeTimeout = timeout
	}}
}

var (
	redisConnections map[string]*Redis = make(map[string]*Redis, 0)
)

func NewRedis(options ...RedisOption) *Redis {
	redis := &Redis{}
	redis.conn = nil
	redis.protocol = "tcp"
	redis.address = "redis:6379"
	redis.connectTimeout = time.Second
	redis.readTimeout = time.Second
	redis.writeTimeout = time.Second
	for _, option := range options {
		option.f(redis)
	}
	return redis
}

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
		r.conn, err = redis.Dial(r.protocol, r.address, connectTimeout, readTimeout, writeTimeout)
	}
	return r.conn, err
}
