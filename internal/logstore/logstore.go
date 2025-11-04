// Package logstore stores the messages and gets messages based on offset
package logstore

import (
	"os"
	"strings"
)

func Write(msg string) {
	f, err := os.OpenFile("data.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(strings.Trim(msg, ""))
	if err != nil {
		panic(err)
	}

	f.Close()
}
