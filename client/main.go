package client

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type EventRelayClient struct {
	dbPath string
}

func NewEventRelayClient(dbPath string) *EventRelayClient {
	log.Printf("Initializing remote client for: %s\n", dbPath)

	return &EventRelayClient{
		dbPath: dbPath,
	}
}

func (c *EventRelayClient) Publish(payload string) {
	log.Printf("Publishing event with payload: %s\n", payload)

	if !json.Valid([]byte(payload)) {
		log.Fatal("[ERROR] Invalid payload, cannot enqueue")
	}

	connStr := fmt.Sprintf("%s?_journal_mode=WAL", c.dbPath)
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		log.Fatalf("[FATAL] CLIENT Could not open the database: %v", err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO events (payload, status) VALUES (?, ?)", payload, "pending")
	if err != nil {
		log.Fatalf("[ERROR] Failed to insert event into local buffer: %v", err)
	}

	id, _ := result.LastInsertId()
	log.Printf("[SUCCESS] Event %d enqueued in SQLite", id)

}
