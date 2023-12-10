package builtins_test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/mgs0196/CSCE4600/Project2/builtins"
)

func TestPrintCallerInfo(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr error
	}{
		{
			name: "test printCallerInfo",
			// Update the expected output based on your expectations
			want:    "ExpectedCallerInfo",
			wantErr: nil,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect standard output to capture printed output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Cleanup to reset standard output
			t.Cleanup(func() {
				os.Stdout = old
			})

			// Run the test
			err := builtins.PrintCallerInfo()

			// Close the write end of the pipe to stop the infinite loop in the capture
			w.Close()

			// Read the captured output from the pipe
			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)

			// Check the printed output
			got := buf.String()
			if got != tt.want {
				t.Errorf("PrintCallerInfo() got = %v, want %v", got, tt.want)
			}

			// Check the error
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("PrintCallerInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
