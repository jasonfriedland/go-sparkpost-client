package main

// Usage:
// 	SPARKPOST_API_KEY=<your API key> \
// 	SPARKPOST_API_URL=<your API URL> \
// 	go run cmd/sp/main.go <from> <to>

import (
	"log"
	"os"

	"github.com/SparkPost/gosparkpost"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	fromAddr = kingpin.Arg("from", "From address.").Required().String()
	toAddr   = kingpin.Arg("to", "To address.").Required().String()
)

func main() {
	kingpin.Parse()

	// Config
	cfg := &gosparkpost.Config{
		BaseUrl:    os.Getenv("SPARKPOST_API_URL"),
		ApiKey:     os.Getenv("SPARKPOST_API_KEY"),
		ApiVersion: 1,
	}

	// Client
	var client gosparkpost.Client
	err := client.Init(cfg)
	if err != nil {
		log.Fatalf("SparkPost client init failed: %s\n", err)
	}

	// Transmission and content
	tx := &gosparkpost.Transmission{
		Recipients: []string{*toAddr},
		Content: gosparkpost.Content{
			HTML:    "<p>This is a test.</p>",
			From:    *fromAddr,
			Subject: "This is a test",
		},
	}

	// Send
	id, _, err := client.Send(tx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transmission sent with id: %s\n", id)
}
