package whois

import (
	"strings"
	"time"
)

func findDate(whois, startString string) (time.Time, error) {
	startIndex := strings.Index(whois, startString) + len(startString)
	endIndex := strings.Index(whois[startIndex:], "\n") + startIndex
	date := whois[startIndex:endIndex]
	date = strings.TrimSpace(date)
	parse, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return parse, err
	}
	return parse, nil
}
