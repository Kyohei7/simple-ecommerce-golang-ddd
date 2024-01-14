package main

import (
	"log"
	"simple-ecommerce/cmd/external/database"
	"simple-ecommerce/cmd/internal/config"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("Database Connected")
	}
}
