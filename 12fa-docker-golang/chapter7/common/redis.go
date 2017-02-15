package common

import "time"
import "github.com/garyburd/redigo/redis"

// Redis connection struct
type Redis struct {
	conn                                      redis.Conn
	address                                   string
	connectTimeout, readTimeout, writeTimeout time.Duration
}

// RedisOption functional options
type RedisOption struct {
	f func(*Redis)
}

// RedisAddress - specify hostname of redis server (including port)
func RedisAddress(address string) RedisOption {
	return RedisOption{func(do *Redis) {
		do.address = address
	}}
}

// RedisConnectTimeout - set connection timeout for the Redis client
func RedisConnectTimeout(timeout time.Duration) RedisOption {
	return RedisOption{func(do *Redis) {
		do.connectTimeout = timeout
	}}
}

// RedisReadTimeout - set a read timeout for a Redis operation
func RedisReadTimeout(timeout time.Duration) RedisOption {
	return RedisOption{func(do *Redis) {
		do.readTimeout = timeout
	}}
}

// RedisWriteTimeout - set a write timeout for a Redis operation
func RedisWriteTimeout(timeout time.Duration) RedisOption {
	return RedisOption{func(do *Redis) {
		do.writeTimeout = timeout
	}}
}

var (
	redisConnections = make(map[string]*Redis, 0)
)

// NewRedis creates a new redis connection
func NewRedis(options ...RedisOption) *Redis {
	redis := &Redis{}
	redis.conn = nil
	redis.address = "redis:6379"
	redis.connectTimeout = time.Second
	redis.readTimeout = time.Second
	redis.writeTimeout = time.Second
	for _, option := range options {
		option.f(redis)
	}
	return redis
}

// Save redis connection for reuse
func (t *Redis) Save(names ...string) {
	connection := "default"
	for _, name := range names {
		connection = name
	}
	redisConnections[connection] = t
}

// GetRedis retrieves a saves redis connection for reuse
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
