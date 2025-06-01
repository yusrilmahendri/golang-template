package main

import (
	"log"
	"os"

	"go-clean-template/config"
	"go-clean-template/internal/app"
	"go-clean-template/pkg/mysql"
)

func main() {
	// 1. Load env/config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// 2. Init MySQL
	mysqlCfg := mysql.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := mysql.NewMysqlConnection(mysqlCfg)
	if err != nil {
		log.Fatalf("DB Error: %v", err)
	}

	// 3. Inject db ke app
	app.Run(cfg, db) // <-- nanti kamu ubah Run() agar menerima db juga
}
