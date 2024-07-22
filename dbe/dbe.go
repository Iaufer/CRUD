package dbe

import (
	"crud/trans"
	"database/sql"
	"fmt"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		info.Host, info.Port, info.Username, info.DBName, info.SSLMode, info.Password))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InsertBook(db *sql.DB, b trans.Book) error {
	if db == nil {
		fmt.Println("Реально нил")
	}
	_, err := db.Exec("insert into books (id, name, price) values ($1, $2, $3)",
		b.ID, b.Name, b.Price)
	return err
}
