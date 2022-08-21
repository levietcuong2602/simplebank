package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/levietcuong2602/simplebank/api"
	db "github.com/levietcuong2602/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_SOURCE := os.Getenv("DB_SOURCE")
	PORT := os.Getenv("PORT")
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("Cannot create Server: ", err)
	}
	err = server.Start(":" + PORT)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
