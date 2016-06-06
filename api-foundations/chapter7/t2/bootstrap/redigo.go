package bootstrap

import "time"
import "github.com/garyburd/redigo/redis"
import "github.com/youtube/vitess/go/pools"
import "golang.org/x/net/context"

type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

var (
	pool        *pools.ResourcePool
	hasPool     = false
	serverIndex = 1
)

func getServerName() string {
	name := "redis:6379"
	serverIndex++
	if serverIndex > 2 {
		serverIndex = 1
	}
	return name
}

func RedigoPool() *pools.ResourcePool {
	if !hasPool {
		capacity := 2    // hold two connections
		maxCapacity := 4 // hold up to 4 connections
		idleTimeout := time.Minute
		pool = pools.NewResourcePool(func() (pools.Resource, error) {
			serverName := getServerName()
			c, err := redis.Dial("tcp", serverName, connectTimeout, readTimeout, writeTimeout)
			return ResourceConn{c}, err
		}, capacity, maxCapacity, idleTimeout)
	}
	return pool
}

func RedigoDo(commandName string, params ...interface{}) (interface{}, error) {
	ctx := context.TODO()
	r, err := pool.Get(ctx)
	if err != nil {
		return "", err
	}
	defer pool.Put(r)

	c := r.(ResourceConn)
	return c.Do(commandName, params...)
}
