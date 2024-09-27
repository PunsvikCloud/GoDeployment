package main

import (
	"strings"
	"testing"
)

func Test_getStringWithTimestamp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test OK prefix",
			want: "OK - ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringWithTimestamp(); !strings.HasPrefix(got, tt.want) {
				t.Errorf("getStringWithTimestamp() = %v, want prefix %v", got, tt.want)
			}
		})
	}
}
