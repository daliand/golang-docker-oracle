package main

import (
	"database/sql"
	"log"

	//_ "github.com/mattn/go-oci8"
	_ "gopkg.in/rana/ora.v4"
)

var (
	db *sql.DB
)

func init() {
	var errOracle error

	db, errOracle = sql.Open("ora", "system/plonka99@//127.0.0.1:1521/xe")

	if errOracle != nil {
		log.Fatal("initOracle sql.Open: " + errOracle.Error())
	}
	errOracle = db.Ping()
	if errOracle != nil {
		log.Fatal("initOracle Ping: " + errOracle.Error())
	}
}

func main() {
	defer db.Close()

	var test string
	row := db.QueryRow("SELECT 'Golang Oracle Test' FROM DUAL")
	err := row.Scan(&test)

	if err != nil {
		log.Printf("Oracle error: %v", err)
	}

	log.Printf("DB Result: %v", test)
}
