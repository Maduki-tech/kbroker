// Package protocol is for the parsing of the messages and the logic
package protocol

import (
	"log"
	"strings"
)

// ParseMessage parses the message raw content
func ParseMessage(msg string) {
	splitString := strings.Split(msg, " ")
	cmd := splitString[0]
	text := strings.Join(splitString[1:], " ")

	switch cmd {
	case "PRODUCE":
		log.Println(text)
	case "FETCH":
		log.Printf("Get %s\n", text)
	default:
		log.Fatalf("%s is not a correct command", cmd)
	}
}
