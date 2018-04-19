package factory

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DatabaseProfiler interface {
	Pre(query string, args ...interface{}) *DatabaseProfilerContext
	Post(*DatabaseProfilerContext)
	Flush()
}

type DatabaseProfilerContext struct {
	Query string
	Args  string
	Time  time.Time
}

type DatabaseCredential struct {
	DSN string
}

type DatabaseFactory struct {
	credentials map[string]*DatabaseCredential
	instances   map[string]*DB

	ProfilerStdout DatabaseProfilerStdout
	ProfilerMemory DatabaseProfilerMemory
}

var Database *DatabaseFactory

func init() {
	Database = &DatabaseFactory{}
	Database.credentials = make(map[string]*DatabaseCredential)
	Database.instances = make(map[string]*DB)
}

func (r *DatabaseFactory) Add(name string, config interface{}) {
	switch val := config.(type) {
	case string:
		r.credentials[name] = &DatabaseCredential{DSN: val}
	case DatabaseCredential:
		r.credentials[name] = &val
	default:
		panic("factory.Database.Add can take config as string|factory.DatabaseCredential")
	}
}

func (r *DatabaseFactory) GetDSN(name string) (string, error) {
	addOption := func(s, match, option string) string {
		if !strings.Contains(s, match) {
			s += option
		}
		return s
	}

	if value, ok := r.credentials[name]; ok {
		value.DSN = addOption(value.DSN, "?", "?")
		value.DSN = addOption(value.DSN, "collation=", "&collation=utf8_general_ci")
		value.DSN = addOption(value.DSN, "parseTime=", "&parseTime=true")
		value.DSN = addOption(value.DSN, "loc=", "&loc=Local")
		value.DSN = strings.Replace(value.DSN, "?&", "?", 1)
		return value.DSN, nil
	}
	return "", fmt.Errorf("No configuration found for database: %v", name)
}

func (r *DatabaseFactory) Get(dbName ...string) (*DB, error) {
	names := dbName
	if len(names) == 0 {
		names = []string{"default"}
	}
	for _, name := range names {
		if value, ok := r.instances[name]; ok {
			return value, nil
		}
		dsn, _ := r.GetDSN(name)
		if dsn != "" {
			handle, err := sqlx.Open("mysql", dsn)
			if err != nil {
				return nil, err
			}
			r.instances[name] = &DB{handle, nil}
			return r.instances[name], nil
		}
	}
	return nil, fmt.Errorf("No configuration found for database: %v", names)
}

// DB struct encapsulates sqlx.DB to add new functions
type DB struct {
	*sqlx.DB

	Profiler DatabaseProfiler
}

func (r *DB) SetFields(fields []string) string {
	idx := 0
	sql := ""
	for _, field := range fields {
		if idx > 0 {
			sql = sql + ", "
		}
		idx++
		sql = sql + field + "=:" + field
	}
	return sql
}

func (r *DB) Select(dest interface{}, query string, args ...interface{}) error {
	var err error
	if r.Profiler != nil {
		ctx := r.Profiler.Pre(query, args...)
		err = r.DB.Select(dest, query, args...)
		r.Profiler.Post(ctx)
	} else {
		err = r.DB.Select(dest, query, args...)
	}
	// clear no rows returned error
	if err == sql.ErrNoRows {
		return nil
	}
	return errors.Wrap(err, "select query failed")
}

func (r *DB) Get(dest interface{}, query string, args ...interface{}) error {
	var err error
	if r.Profiler != nil {
		ctx := r.Profiler.Pre(query, args...)
		err = r.DB.Get(dest, query, args...)
		r.Profiler.Post(ctx)
	} else {
		err = r.DB.Get(dest, query, args...)
	}
	// clear no rows returned error
	if err == sql.ErrNoRows {
		return nil
	}
	return errors.Wrap(err, "get query failed")
}

func (r *DB) set(data interface{}) string {
	message_value := reflect.ValueOf(data)
	if message_value.Kind() == reflect.Ptr {
		message_value = message_value.Elem()
	}

	message_fields := make([]string, message_value.NumField())

	for i := 0; i < len(message_fields); i++ {
		fieldType := message_value.Type().Field(i)
		message_fields[i] = fieldType.Tag.Get("db")
	}

	sql := ""
	for _, tagFull := range message_fields {
		if tagFull != "" && tagFull != "-" {
			tag := strings.Split(tagFull, ",")
			sql = sql + " " + tag[0] + "=:" + tag[0] + ","
		}
	}
	return sql[1 : len(sql)-1]
}

func (r *DB) Replace(table string, data interface{}) error {
	sql := "replace into " + table + " set " + r.set(data)
	_, err := r.NamedExec(sql, data)
	return err
}

func (r *DB) Insert(table string, data interface{}) error {
	sql := "insert into " + table + " set " + r.set(data)
	_, err := r.NamedExec(sql, data)
	return err
}

func (r *DB) InsertIgnore(table string, data interface{}) error {
	sql := "insert ignore into " + table + " set " + r.set(data)
	_, err := r.NamedExec(sql, data)
	return err
}
