package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func NewConnection() {
	host := os.Getenv("MYSQL_HOST")
	login := os.Getenv("MYSQL_LOGIN")
	password := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_NAME")
	port := os.Getenv("MYSQL_PORT")

	dataSourceName := fmt.Sprint(login, ":", password, "@tcp(", host, ":", port, ")/", name)
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
}
