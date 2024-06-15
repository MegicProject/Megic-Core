package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@Talenta123@tcp(127.0.0.1:3306)/cenel_db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Connected to MySQL database!")
	}

	http.HandleFunc("/", handler)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
