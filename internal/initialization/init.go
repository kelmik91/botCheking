package initialization

import (
	"botCheking/internal/db/mysql"
	"github.com/joho/godotenv"
)

func Initialization() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file" + err.Error())
	}
	mysql.NewConnection()
}
