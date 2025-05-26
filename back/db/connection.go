package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func Connect() {
	once.Do(func() {
		connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=cinema_db sslmode=disable"
		var err error

		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Ошибка при подключении к БД:", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatal("БД не отвечает:", err)
		}

		fmt.Println("✅ Подключено к PostgreSQL")
	})
}

func GetDB() *sql.DB {
	return db
}
