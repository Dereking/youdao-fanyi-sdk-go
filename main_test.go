package main

import (
	"testing"
)

func Test_translateToCn(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"“It’s either traffic forever or tunnels.”"},
		{"a boy"},
		{"I'm a good boy."},
		{"Are you ok?"},
		{"Are you ok? I'm a good boy."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := translateToCn(tt.name)
			t.Log(a, err)
		})

	}
}
