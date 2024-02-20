package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Явный импорт драйвера MySQL
)

const (
	host     = "localhost"
	port     = 3306
	user     = "yourusername"
	password = ""
	dbname   = "internet-shop"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	orderNumbers := os.Args[1:]
}
