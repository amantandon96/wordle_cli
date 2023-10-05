package config

import (
	"testing"
)

func Test_initializeConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"config test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println("init complete")
		})
	}
}
