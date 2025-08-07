package main

import (
	"database/sql"
	"jwt/auth"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	 _ "github.com/go-sql-driver/mysql"
)

func main(){
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env")
    }
	dsn := os.Getenv("DB_DSN")
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
	  if err := db.Ping(); err != nil {
        log.Fatal("DB ulanishda xatolik:", err)
    }
	  auth.DB = db

	      http.HandleFunc("/register", auth.RegisterHandler)
    http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/refresh", auth.RefreshHandler)

    log.Println("Server is running at :8002")
    http.ListenAndServe(":8002", nil)
}