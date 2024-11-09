package main

import (
	"botCheking/internal/db/mysql"
	"botCheking/internal/db/mysql/repository/sites"
	"botCheking/internal/whois"
	"fmt"
	"github.com/joho/godotenv"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file" + err.Error())
	}
	mysql.NewConnection()
}

func main() {
	hosts := sites.Sites()
	if len(hosts) < 1 {
		return
	}

	ch := make(chan whois.Site, 1)
	go whois.Check(hosts, ch)

	for h := range ch {
		if h.Date.Before(time.Now().Add(7 * 24 * time.Hour)) {
			fmt.Printf("Аренда домена %s заканчивается %s\n", h.URL, h.Date.Format("02/01/2006 15:04"))
		}
	}
}
