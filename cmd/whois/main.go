package main

import (
	"botCheking/internal/db/mysql/repository/sites"
	"botCheking/internal/initialization"
	"botCheking/internal/whois"
	"fmt"
	"time"
)

func init() {
	initialization.Initialization()
}

func main() {
	hosts := sites.Sites()
	if len(hosts) < 1 {
		return
	}

	ch := make(chan whois.Site, 1)
	go whois.GetInfo(hosts, ch)

	for h := range ch {
		if h.Date.Before(time.Now().Add(7 * 24 * time.Hour)) {
			fmt.Printf("Аренда домена %s заканчивается %s\n", h.URL, h.Date.Format("02/01/2006 15:04"))
		}
	}
}
