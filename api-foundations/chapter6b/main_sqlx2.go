package main

import "log"
import "foundations/bootstrap"
import "github.com/davecgh/go-spew/spew"

type Database struct {
	Name string `db:"Database"`
}

func main() {
	pool := bootstrap.SqlxConnectionPool()
	defer pool.Close()

	db, err := bootstrap.SqlxGetConnection()
	if err != nil {
		log.Fatal("Error when connecting: ", err)
	}
	defer bootstrap.SqlxReleaseConnection(db)

	databases := []Database{}
	err = db.Select(&databases, "show databases")
	if err != nil {
		log.Fatal("Error in query: ", err)
	}

	spew.Dump(databases)
}
