package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlitePath := flag.String("sqlite", "", "Path to the SQLite3 file") // -flag name, default val, help
	flag.Parse()

	if *sqlitePath == "" {
		log.Fatal("Error: Please specify the path to the SQLite3 file using -sqlite")
	}

	db, err := sql.Open("sqlite3", *sqlitePath)
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}
	defer db.Close()

	query := `SELECT * FROM trace ORDER BY start_time ASC LIMIT 1`
	row := db.QueryRow(query)

	var taskID, parentID, kind, what, location string
	var startTime, endTime float64


	err = row.Scan(&taskID, &parentID, &kind, &what, &location, &startTime, &endTime)
	if err != nil {
		log.Fatalf("Error fetching row: %v", err)
	}

	fmt.Printf("task_id: %s parent_id: %s kind: %s what: %s location: %s\nstart_time: %.12f(s) end_time: %.12f(s)\n", taskID, parentID, kind, what, location, startTime, endTime)
}