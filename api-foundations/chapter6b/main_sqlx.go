package main

import "log"
import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"
import "github.com/davecgh/go-spew/spew"

type Database struct {
	Name string `db:"Database"`;
}

func main() {
	db, err := sqlx.Open("mysql", "api:api@tcp(db1:3306)/api");
	if err != nil {
		log.Fatal("Error when connecting: ", err);
	}
	databases := []Database{};
	err = db.Select(&databases, "show databases");
	if err != nil {
		log.Fatal("Error in query: ", err);
	}	
	spew.Dump(databases);
}