package main

import (
	"github.com/markitos-it/mdk-event-relay-client/client"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	runClient()
}

func runClient() {
	dbPath := "../mdk-event-relay/events.db"
	client := client.NewEventRelayClient(dbPath)
	client.Publish(`{
		"name": "markitos",
		"date": "any",
		"send": true
	}`)

}
