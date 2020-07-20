package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "mydb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port = %d user = %s"+
		"password = %s dbname = %s sslmode = disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// fmt.Println("Hi")
	// err = db.Ping()
	// if err != nil {
	// panic(err)
	// }

	//DB = db

	fmt.Println("Successfully connected!")
}
