package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	tests := []struct {
		name       string
		readstring string
		want       string
		wantErr    error
	}{
		{name: "正常", readstring: "abc", want: "abc"},
		{name: "empty", readstring: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := cat(buf, strings.NewReader((tt.readstring)))
			got := buf.String()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err= %v, want %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("cat(&bytes.Buffer{}, %s) = %s, want %s", tt.readstring, got, tt.want)
			}
		})
	}
}
