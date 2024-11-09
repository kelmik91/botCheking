package sites

import (
	"botCheking/internal/db/mysql"
	"log"
)

func Sites() []string {
	rows, err := mysql.DB.Query("SELECT name FROM domains WHERE is_active = 1")
	if err != nil {
		log.Println("Ошибка при получении строк из БД", err)
	}
	defer rows.Close()

	var hosts []string
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Println("Ошибка при сканировании строки из БД", err)
		}
		hosts = append(hosts, s)
	}

	return hosts
}
