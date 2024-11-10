package whois

import (
	"log"
	"strings"
	"time"
)

type Site struct {
	URL  string
	Date time.Time
}

func GetInfo(hosts []string, ch chan<- Site) {
	defer close(ch)

	for _, host := range hosts {
		w, err := Whois(host)
		if err != nil {
			log.Fatal(err)
			return
		}

		date := time.Time{}
		if strings.Contains(w, "paid-till:") {
			startString := "paid-till:"
			date, err = findDate(w, startString)
			if err != nil {
				panic(err)
			}
		}
		if strings.Contains(w, "Registry Expiry Date:") {
			startString := "Registry Expiry Date:"
			date, err = findDate(w, startString)
			if err != nil {
				panic(err)
			}
		}
		if strings.Contains(w, "Registrar Registration Expiration Date:") {
			startString := "Registrar Registration Expiration Date:"
			date, err = findDate(w, startString)
			if err != nil {
				panic(err)
			}
		}

		loc, _ := time.LoadLocation("Europe/Moscow")
		dateEnd := date.In(loc)

		ch <- Site{
			URL:  host,
			Date: dateEnd,
		}
	}
}
