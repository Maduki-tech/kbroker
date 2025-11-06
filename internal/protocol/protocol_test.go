package protocol_test

import (
	"testing"

	"github.com/maduki-tech/kbroker/internal/protocol"
)

func TestParseMessage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		msg     string
		want    string
		wantErr bool
	}{
		{
			name:    "Successfull Produce",
			msg:     "PRODUCE mock_success",
			want:    "mock_success",
			wantErr: false,
		},
		{
			name:    "Successfull FETCH",
			msg:     "FETCH fetch_success",
			want:    "fetch_success",
			wantErr: false,
		},
		{
			name:    "Successfull Error",
			msg:     "fetch_success",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := protocol.ParseMessage(tt.msg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseMessage() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseMessage() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("ParseMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
