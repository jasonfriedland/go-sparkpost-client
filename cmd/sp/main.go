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
	fromAddr   = kingpin.Arg("from", "From email address.").Required().String()
	toAddr     = kingpin.Arg("to", "To email address.").Required().String()
	returnPath = kingpin.Flag("return-path", "Return path address.").Short('r').Default(*fromAddr).String()
	subject    = kingpin.Flag("subject", "Email subject.").Short('s').String()
)

// OS env vars.
var (
	spBaseURL = os.Getenv("SPARKPOST_API_URL")
	spAPIKey  = os.Getenv("SPARKPOST_API_KEY")
)

// getStdIn gets the emnail input from stdin, or sets a default if empty.
func getStdIn(defaultInput string) string {
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
		input = strings.Join(lines, "\n")
	} else {
		input = defaultInput
	}

	return strings.Trim(input, "\n") + "\n" // append our chomped final newline
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
		BaseUrl:    spBaseURL,
		ApiKey:     spAPIKey,
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
		ReturnPath: *returnPath,
		Content: gosparkpost.Content{
			Text:    getStdIn("No email content provided."),
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
