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

func TestUrlify(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"Mr John Smith", "Mr%20John%20Smith"},
		{"MrJohnSmith", "MrJohnSmith"},
		{"MrJohnSmith ", "MrJohnSmith%20"},
		{" MrJohnSmith", "%20MrJohnSmith"},
		{" MrJohnSmith ", "%20MrJohnSmith%20"},
	}

	for _, c := range cases {
		got := urlify(c.s)

		require.Equal(t, got, c.want, "got %v, want %v", got, c.want)
	}
}

func TestPalindromePermutation(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"Tact Coa", true},
		{"ttt", true},
		{"tttaaaccc", true},
		{"t t t a a a c c c", true},
		{"xyz", false},
		{"xyz abc", false},
	}

	for _, c := range cases {
		got := palindromePermutation(c.s)

		require.Equal(t, got, c.want, "case: %v\ngot %v, want %v", c.s, got, c.want)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"ttt", true},
		{"taco cat", true},
		{"atco cta", true},
		{"tact coa", false},
		{"tttaaaccc", false},
		{"t t t a a a c c c", false},
		{"xyz", false},
		{"xyz%!*!%zyx", true},
		{"xyz abc", false},
	}

	for _, c := range cases {
		got := isPalindrome(c.s)

		require.Equal(t, got, c.want, "case: %v\ngot %v, want %v", c.s, got, c.want)
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"Mr John Smith", "htimS nhoJ rM"},
	}

	for _, c := range cases {
		got := reverse(c.s)

		require.Equal(t, got, c.want, "got %v, want %v", got, c.want)
	}
}

func TestMergeSortedArrays(t *testing.T) {
	cases := []struct {
		a, b []int
		want []int
	}{
		{nil, nil, []int{}},
		{[]int{}, nil, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{4, 6, 30}, []int{4, 6, 30}},
		{[]int{0, 3, 4, 31}, []int{4, 6, 30}, []int{0, 3, 4, 4, 6, 30, 31}},
	}

	for _, c := range cases {
		got := mergeSortedArrays(c.a, c.b)

		require.Equal(t, c.want, got, "got %v, want %v", got, c.want)
	}
}
