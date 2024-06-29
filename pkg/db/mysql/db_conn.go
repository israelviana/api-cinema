package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"sync"
)

var (
	instance *sql.DB
	once     sync.Once
)

func NewMySQLConnection() *sql.DB {
	once.Do(func() {
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}

		err = db.Ping()
		if err != nil {
			panic(err.Error())
		}

		instance = db
	})

	return instance
}
