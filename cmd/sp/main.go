package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/SparkPost/gosparkpost"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Kingpin args/flags.
var (
	fromAddr = kingpin.Arg("from", "From email address.").Required().String()
	toAddr   = kingpin.Arg("to", "To email address.").Required().String()
	subject  = kingpin.Flag("subject", "Email subject.").Short('s').String()
)

// getInput gets the emnail input from stdin, or sets a default if empty.
func getInput() string {
	var err error
	var lines []string
	var input string

	// Check for piped in input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Body input is being piped in
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			log.Println("error reading stdin:", err)
		}
		input = strings.Join(lines, "\n") + "\n" // append our chomped final newline
	} else {
		input = "No email content provided."
	}

	return input
}

// Main.
func main() {
	var err error
	var client gosparkpost.Client
	var tx *gosparkpost.Transmission

	// Parse args
	kingpin.Parse()

	// Config
	cfg := &gosparkpost.Config{
		BaseUrl:    os.Getenv("SPARKPOST_API_URL"),
		ApiKey:     os.Getenv("SPARKPOST_API_KEY"),
		ApiVersion: 1,
	}

	// Client
	err = client.Init(cfg)
	if err != nil {
		log.Fatalln("client init failed:", err)
	}

	// Transmission and content
	tx = &gosparkpost.Transmission{
		Recipients: []string{*toAddr},
		ReturnPath: *fromAddr,
		Content: gosparkpost.Content{
			Text:    getInput(),
			From:    *fromAddr,
			Subject: *subject,
		},
	}

	// Send
	id, _, err := client.Send(tx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("sent with tx id:", id)
}
