package data

import (
	"database/sql"
	"fmt"
	"log"
	"url_shortener/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {

	connString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.DBUSER, config.DBPASS, config.DBNAME)

	db, err := sql.Open(config.DBDRIVER, connString)

	if err != nil {
		log.Fatalln(err.Error())
		defer db.Close()
	}
	return db
}
