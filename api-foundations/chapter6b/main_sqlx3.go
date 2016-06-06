package main

import "log"
import "fmt"
import "foundations/bootstrap"

type SleepResult struct {
	Result string `db:"sleepfor"`
}

func main() {
	pool := bootstrap.SqlxConnectionPool()
	defer pool.Close()

	// warm up the connection pool
	for i := 0; i < 5; i++ {
		db, _ := bootstrap.SqlxGetConnection()
		db.Ping()
		bootstrap.SqlxReleaseConnection(db)
	}

	fmt.Printf("[%.4f] Start\n", bootstrap.Now())

	sleep1_chan := make(chan string, 1)
	sleep2_chan := make(chan string, 1)

	go func() {
		db, err := bootstrap.SqlxGetConnection()
		if err != nil {
			log.Fatal("Error when connecting: ", err)
		}
		defer bootstrap.SqlxReleaseConnection(db)

		fmt.Printf("[%.4f] Run sleep 100ms\n", bootstrap.Now())

		fromSleep := SleepResult{}
		err = db.Get(&fromSleep, "select sleep(0.1) as sleepfor")
		if err != nil {
			sleep1_chan <- "ERROR"
			return
		}
		sleep1_chan <- fromSleep.Result
	}()

	go func() {
		db, err := bootstrap.SqlxGetConnection()
		if err != nil {
			log.Fatal("Error when connecting: ", err)
		}
		defer bootstrap.SqlxReleaseConnection(db)

		fmt.Printf("[%.4f] Run sleep 200ms\n", bootstrap.Now())

		fromSleep := SleepResult{}
		err = db.Get(&fromSleep, "select sleep(0.2) as sleepfor")
		if err != nil {
			sleep2_chan <- "ERROR"
		}
		sleep2_chan <- fromSleep.Result
	}()

	var result string
	result = <-sleep1_chan
	fmt.Printf("[%.4f] End Sleep 100ms, result %s\n", bootstrap.Now(), result)
	result = <-sleep2_chan
	fmt.Printf("[%.4f] End Sleep 200ms, result %s\n", bootstrap.Now(), result)
}
