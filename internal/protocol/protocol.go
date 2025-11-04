// Package protocol is for the parsing of the messages and the logic
package protocol

import (
	"fmt"
	"strings"
)

// ParseMessage parses the message raw content
func ParseMessage(msg string) (string, error) {
	splitString := strings.Split(msg, " ")
	cmd := splitString[0]
	text := strings.Join(splitString[1:], " ")

	switch cmd {
	case "PRODUCE":
		// TODO: Add to log files
		return text, nil
	case "FETCH":
		// TODO: get from log file
		return text, nil
	default:
		return "", fmt.Errorf("%s is not a correct command", cmd)
	}
}
