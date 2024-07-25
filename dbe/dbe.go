package dbe

import (
	"crud/trans"
	"database/sql"
	"fmt"
	"log"
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

func GetAllBooks(db *sql.DB) []trans.Book {
	rows, err := db.Query("select * from books")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	books := make([]trans.Book, 0)

	for rows.Next() {
		b := trans.Book{}

		err := rows.Scan(&b.ID, &b.Name, &b.Price)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, b)
	}
	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

	return books
}

func DeleteBook(db *sql.DB, id int) error {
	_, err := db.Exec("delete from books where id = $1", id)
	return err
}
