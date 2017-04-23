package service

import (
	"strings"
	// import database driver for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Database connection factory object
type Database struct {
	conn *sqlx.DB
	err  error

	dsn *string
}

// Interface for used flag functions
type databaseFlags interface {
	String(name, value, usage string) *string
}

// Flags registers a new parameter in a compatible flags package
func (r *Database) Flags(name, usage string, flag databaseFlags) {
	r.dsn = flag.String(name, "api:api@tcp(db1:3306)/api", usage)
}

// Get creates and returns a new database connection
func (r *Database) Get() (*sqlx.DB, error) {
	dsn := *r.dsn
	if !strings.Contains(dsn, "?") {
		dsn = dsn + "?"
	}
	// set some default values for dsn config
	defaults := map[string]string{
		"collation": "utf8_general_ci",
		"parseTime": "true",
		"loc":       "Local",
	}
	// append default values
	for key, value := range defaults {
		if !strings.Contains(dsn, key+"=") {
			dsn = dsn + "&" + key + "=" + value
		}
	}
	dsn = strings.Replace(dsn, "?&", "?", 1)
	r.conn, r.err = sqlx.Open("mysql", dsn)
	return r.conn, r.err
}

// Close releases an existing database connection
func (r *Database) Close() {
	r.conn.Close()
}
