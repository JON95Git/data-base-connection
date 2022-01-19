package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //implicit import
)

func main() {
	// Data base credentials and other configurations
	stringConnection := "golang:golang@/godatabase?charset=utf8&parseTime=True&loc=Local"
	dataBase, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}

	defer dataBase.Close()

	// Try to ping data base connection
	if err = dataBase.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open connection")

	lines, err := dataBase.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()
	fmt.Println(lines)
}
