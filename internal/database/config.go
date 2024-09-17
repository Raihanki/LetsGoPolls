package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
		return nil
	}

	log.Println("database connected")
	return db
}
