package logstore_test

import (
	"testing"

	"github.com/maduki-tech/kbroker/internal/logstore"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		msg string
	}{
		{
			name: "Basic Hello",
			msg:  "Hello 1",
		},
		{
			name: "Basic Hello",
			msg:  "Hello 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logstore.Write(tt.msg)
		})
	}
}
