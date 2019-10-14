package cci

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueStr(t *testing.T) {
	cases := map[string]bool{
		"asdfghlkqwe":     true,
		"asdfghlkqwe1234": true,
		"asrfghlkqre1234": false,
	}

	for k, v := range cases {
		got := uniqueStr(k)
		got2 := uniqueStr2(k)

		require.Equal(t, got, v, "%v: got %v, want %v", k, got, v)
		require.Equal(t, got2, v, "%v: got %v, want %v", k, got, v)
	}
}

func TestCheckPerm(t *testing.T) {
	cases := []struct {
		s    string
		t    string
		want bool
	}{
		{"a", "asdfghlkqwe", true},
		{"dfgh", "asdfghlkqwe", true},
		{"asdfghlkqwe", "asdfghlkqwe", true},
		{"asd", "asdfghlkqwe", true},
		{"qwe", "asdfghlkqwe", true},

		{"qweg", "asdfghlkqwe", false},
		{"kl", "asdfghlkqwe", false},
		{"z", "asdfghlkqwe", false},
		{"asdfghlkqwed", "asdfghlkqwe", false},
	}

	for _, c := range cases {
		got := checkPerm(c.s, c.t)

		require.Equal(t, got, c.want, "%v --> %v:\n got %v, want %v", c.s, c.t, got, c.want)
	}
}
