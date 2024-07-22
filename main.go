package main

import (
	"crud/db"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	db, err := db.NewPostgresConnection(db.ConnectionInfo{
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

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("SERVER STARTED")

	http.HandleFunc("/", handle)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
