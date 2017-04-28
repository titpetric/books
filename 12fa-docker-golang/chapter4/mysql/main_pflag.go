package main

import (
	"log"
	"fmt"
	"os"

	"app/service"

	flag "github.com/spf13/pflag"
	"github.com/jmoiron/sqlx"
)

type databaseName struct {
	Name string `db:"Database"`
}

func listDatabases(db *sqlx.DB) ([]databaseName, error) {
	result := []databaseName{}
	err := db.Select(&result, "show databases")
	return result, err
}

func main() {
	flags := flag.NewFlagSet("default", flag.ContinueOnError)

	// this is just a factory struct
	dbFactory := &service.Database{}
	dbFactory.Flags("dsn", "DSN for database connection", flags)

	flags.Parse(os.Args[1:])

	db, err := dbFactory.Get();
	if err != nil {
		log.Println("Error when connecting:", err)
		return
	}

	fmt.Println("Listing databases")
	databases, err := listDatabases(db)
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
