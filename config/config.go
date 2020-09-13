package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT     = 0
	DBDRIVER = ""
	DBPASS   = ""
	DBNAME   = ""
	DBUSER   = ""
	DBURL    = ""
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9090
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBUSER = os.Getenv("DB_USER")
	DBPASS = os.Getenv("DB_PASS")
	DBNAME = os.Getenv("DB_NAME")
}
