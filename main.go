package main

import (
	"crud/dbe"
	"crud/trans"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB // Глобальная переменная для хранения соединения с базой данных

// Инициализация соединения с базой данных
func initDB() {
	var err error
	db, err = dbe.NewPostgresConnection(dbe.ConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "2505",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")
}

// Обработчик HTTP запросов
func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Обработка GET запросов, если нужно
	case http.MethodPost:
		b, err := trans.AddBook(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = dbe.InsertBook(db, b)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to insert book", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	initDB() // Инициализация базы данных при запуске приложения

	fmt.Println("Server started")

	http.HandleFunc("/", Handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
