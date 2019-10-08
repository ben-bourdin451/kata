package cci

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

		assert.Equal(t, got, v, "%v: got %v, want %v", k, got, v)
		assert.Equal(t, got2, v, "%v: got %v, want %v", k, got, v)
	}
}
