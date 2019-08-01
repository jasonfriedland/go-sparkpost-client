package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_getStdIn(t *testing.T) {
	// Test args
	type args struct {
		defaultInput string
	}
	tests := []struct {
		name  string
		stdIn string
		args  args
		want  string
	}{
		// Test cases
		{"Pass stdin, simple case", "x", args{"y"}, "x\n"},
		{"Pass stdin, multiple newlines", "x\n\n\n", args{"y"}, "x\n"},
		{"No stdin, simple case", "", args{"y"}, "y\n"},
		{"No stdin, multiple newlines", "", args{"y\n\n\n"}, "y\n"},
	}

	for _, tt := range tests {
		// Back up STDIN
		oldStdin := os.Stdin
		if tt.stdIn != "" {
			tmp := makeTempFile(tt.stdIn)
			defer os.Remove(tmp.Name())
			os.Stdin = tmp
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := getStdIn(tt.args.defaultInput); got != tt.want {
				t.Errorf("getStdIn() = %v, want %v", got, tt.want)
			}
		})
		// Restore STDIN
		os.Stdin = oldStdin
	}
}

func makeTempFile(input string) *os.File {
	content := []byte(input)
	tmpFile, err := ioutil.TempFile("", "main_test")
	if err != nil {
		return nil
	}
	if _, err := tmpFile.Write(content); err != nil {
		return nil
	}
	if _, err := tmpFile.Seek(0, 0); err != nil {
		return nil
	}
	return tmpFile
}
