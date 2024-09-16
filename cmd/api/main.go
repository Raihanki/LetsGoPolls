package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Raihanki/LetsGoPolls/internal/database"
	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	DB *sql.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading configuration file: %v", err)
	}

	db := database.ConnectDB()
	appCfg := &ApplicationConfig{
		DB: db,
	}

	server := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: appCfg.Routes(),
	}

	log.Println("server start on port :", os.Getenv("APP_PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("error while serving to port 8000: %v", err)
	}
}
