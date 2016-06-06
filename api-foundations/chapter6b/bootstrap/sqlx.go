package bootstrap

import "fmt"
import "time"

import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"

import "github.com/youtube/vitess/go/pools"
import "golang.org/x/net/context"

type SqlxResourceConn struct {
	*sqlx.DB
}

func (r SqlxResourceConn) Close() {
	r.DB.Close()
}

var (
	pool            *pools.ResourcePool
	hasPool         = false
	connectionIndex = 1
)

func SqlxConnectionPool() *pools.ResourcePool {
	if !hasPool {
		capacity := 2    // hold two connections
		maxCapacity := 4 // hold up to 4 connections
		idleTimeout := time.Minute
		pool = pools.NewResourcePool(func() (pools.Resource, error) {
			db, err := sqlx.Open("mysql", "api:api@tcp(db1:3306)/api")
			fmt.Printf("New mysql connection: %d\n", connectionIndex)
			connectionIndex++
			return SqlxResourceConn{db}, err
		}, capacity, maxCapacity, idleTimeout)
	}
	return pool
}

func SqlxGetConnection() (SqlxResourceConn, error) {
	ctx := context.TODO()
	db, err := pool.Get(ctx)
	return db.(SqlxResourceConn), err
}

func SqlxReleaseConnection(r SqlxResourceConn) {
	pool.Put(r)
}
