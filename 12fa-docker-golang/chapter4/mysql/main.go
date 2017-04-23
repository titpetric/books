package main

import (
	"flag"
	"log"
	"fmt"
	"os"

	"app/service"
)

type databaseName struct {
	Name string `db:"Database"`
}

func listDatabases(db *service.Database) ([]databaseName, error) {
	result := []databaseName{}
	conn, err := db.Get()
	if err != nil {
		return result, err
	}
	defer db.Close()

	err = conn.Select(&result, "show databases")
	return result, err
}

func main() {
	flags := flag.NewFlagSet("default", flag.ContinueOnError)

	// this is just a factory struct
	database := &service.Database{}
	database.Flags("dsn", "DSN for database connection", flags)

	flags.Parse(os.Args[1:])

	fmt.Println("Listing databases")
	databases, err := listDatabases(database)
	if err != nil {
		log.Println("Error when listing databases:", err)
		return
	}
	if len(databases) == 0 {
		fmt.Println("No databases exist")
	} else {
		for _, db := range databases {
			fmt.Println(">", db.Name)
		}
	}
}
