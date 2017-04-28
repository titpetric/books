package service

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
	dsn []*string
	serverIndex int
}

// Interface for used flag functions
type redisFlags interface {
	String(name, value, usage string) *string
}

// Flags registers a new parameter in a compatible flags package
func (r *Redis) Flags(name, value, usage string, flag redisFlags) {
	r.dsn = append(r.dsn, flag.String(name, value, usage))
}

// Get a single redis server name (based on flags registrations)
func (r *Redis) getServerName() string {
	name := *r.dsn[r.serverIndex];
	r.serverIndex++
	if r.serverIndex >= len(r.dsn) {
		r.serverIndex = 0
	}
	return name
}

// Returns a connection which needs to be closed
func (r *Redis) Get() redis.Conn {
	if r.pool == nil {
		r.pool = &redis.Pool {
			MaxIdle: 3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", r.getServerName());
			},
		};
	}
	return r.pool.Get()
}

// Execute a command against a temporary pool connection
func (r *Redis) Do(commandName string, params ...interface{}) (interface{}, error) {
	conn := r.Get()
	defer conn.Close()
	return conn.Do(commandName, params...)
}
