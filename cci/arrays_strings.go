package cci

import (
	"strings"
)

// 1.1
// Implement an algorithm to determine if a string has all unique characters.
func uniqueStr(in string) bool {
	m := make(map[rune]struct{})

	for _, c := range in {
		if _, found := m[c]; found {
			return false
		}
		m[c] = struct{}{}
	}

	return true
}

// What if you cannot use additional data structures?
// What is the time complexity? O(n^2)
func uniqueStr2(in string) bool {
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in[i+1:]); j++ {
			if in[i] == in[j] {
				return false
			}
		}
	}
	return true
}

// 1.2
// Given two strings, write a method to decide if one is a permutation of the other.
func checkPerm(s, t string) bool {
	indexMatch := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[indexMatch] {
			indexMatch++
		} else {
			indexMatch = 0
		}

		if indexMatch >= len(s) {
			return true
		}
	}

	return false
}

// 1.3 URLify
// Write a method to replace all spaces in a string with '%20: You may assume that
// the string has sufficient space at the end to hold the additional characters,
// and that you are given the "true" length of the string.
// Note: If implementing in Java, please use a character array so that you can perform
// this operation in place.
func urlify(s string) string {
	var r strings.Builder

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			r.WriteString("%20")
		} else {
			r.WriteString(string(s[i]))
		}
	}
	return r.String()
}

// 1.4 Palindrome Permutation
// Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters. The palindrome does not need to be
// limited to just dictionary words.
func palindromePermutation(s string) bool {
	// for all permutations
	for i := 0; i < len(s); i++ {

	}
	//   if palindrome return true

	return false
}

func isPalindrome(s string) bool {
	fi, bi := 0, len(s)-1

	for fi < bi {
		if string(s[fi]) == " " {
			fi++
			continue
		}

		if string(s[bi]) == " " {
			bi--
			continue
		}

		if strings.ToLower(string(s[fi])) != strings.ToLower(string(s[bi])) {
			return false
		}

		fi++
		bi--
	}

	return true
}
