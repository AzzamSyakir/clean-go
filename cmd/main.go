package main

import (
    "os"
)

func main() {
    // Check command line arguments
    if len(os.Args) > 1 {
        // If the argument is "migration", run migration
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
        // If the argument is "migrate", run migration
        if os.Args[1] == "migrate" {
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

    // If there are no arguments, run the server
    fmt.Println("Server started on port 9000")
    router.RunServer()
}
