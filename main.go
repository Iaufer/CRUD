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

var db *sql.DB

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

func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		g := dbe.GetAllBooks(db)
		for _, value := range g {
			fmt.Println(value.ID, value.Name, value.Price)
		}
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

	case http.MethodDelete:
		id := 0
		fmt.Println("Введите id книги: ")
		fmt.Scanln(&id)

		err := dbe.DeleteBook(db, id)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	initDB()

	fmt.Println("Server started")

	http.HandleFunc("/", Handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
