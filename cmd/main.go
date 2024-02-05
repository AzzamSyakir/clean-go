package main

import (
	"clean-go/internal/config"
	"clean-go/internal/router"
	"clean-go/migration"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Cek argumen command line
	if len(os.Args) > 1 {
		// Jika argumen adalah "migration", jalankan migrasi
		if os.Args[1] == "migration" {
			// Initialize database
			db, err := config.InitDB()
			if err != nil {
				log.Fatal("Error connecting to database:", err)
			}
			err = migration.MigrationDb(db)
			if err != nil {
				log.Fatal("Error running migrations:", err)
			}
			fmt.Println("Migrations successfully!")
			return
		}
	}

	// Jika tidak ada argumen, jalankan servm per
	fmt.Println("Server started on port 9000")
	router.RunServer()
}
