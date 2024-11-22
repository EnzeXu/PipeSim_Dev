package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Command-line argument for the SQLite file path
	sqlitePath := flag.String("sqlite", "", "Path to the SQLite3 file") // -flag name, default val, help
	flag.Parse()

	// Validate input
	if *sqlitePath == "" {
		log.Fatal("Error: Please specify the path to the SQLite3 file using -sqlite")
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite3", *sqlitePath)
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}
	defer db.Close()

	// Query for the entity with the minimum start_time
	query := `SELECT * FROM trace ORDER BY start_time ASC LIMIT 1` // Replace `table_name` with your actual table name
	row := db.QueryRow(query)

	// Fetch all columns of the row (modify according to your schema)
	var taskID, parentID, kind, what, location string
	var startTime, endTime float64


	err = row.Scan(&taskID, &parentID, &kind, &what, &location, &startTime, &endTime) // Adjust based on your table schema
	if err != nil {
		log.Fatalf("Error fetching row: %v", err)
	}

	// Print the entire row
	fmt.Printf("task_id: %s parent_id: %s kind: %s what: %s location: %s\nstart_time: %.12f(s) end_time: %.12f(s)\n", taskID, parentID, kind, what, location, startTime, endTime)
}