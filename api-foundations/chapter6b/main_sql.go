package main

import "log"
import "fmt"
import "github.com/davecgh/go-spew/spew"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func showDatabases(conn *sql.DB, sql string) error {
	stmt, err := conn.Query(sql);
	if err != nil {
		return err;
	}
	defer stmt.Close();
	for stmt.Next() {
		var name string;
		if err := stmt.Scan(&name); err != nil {
			log.Fatal(err);
		}
		fmt.Printf("Database: %s\n", name);
	}
	return nil;
}

func main() {
	db, err := sql.Open("mysql", "api:api@tcp(db1:3306)/api");
	if err != nil {
		log.Fatal("Error when connecting: ", err);
	}

	err = showDatabases(db, "show databases where `database` REGEXP '^inf'");
	if err != nil {
		log.Fatal("Error in query: ", err);
	}
	
	spew.Dump("db");
}