package main

import (
    "fmt"
    "log"

   "github.com/Save-Cash/Save-Cash-pkg/db"
)

func main() {
    if err := db.InitDB(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.CloseDB()

    fmt.Println("Database connection successful!")
}
