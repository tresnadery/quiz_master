package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	var err error
	db, err := sql.Open(os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}
	return db
}
