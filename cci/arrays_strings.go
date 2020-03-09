package cci

import (
	"math"
	"strconv"
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
	ccount := make(map[rune]int)
	for _, c := range s {
		if c != ' ' {
			ccount[c]++
		}
	}

	// more than 1 odd char means this cannot be a palindrome
	odd := 0
	for _, v := range ccount {
		if v%2 != 0 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}

	return true
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

// 1.5 One Away
// There are three types of edits that can be performed on strings:
// insert a character, remove a character, or replace a character.
// Given two strings, write a function to check if they are one edit (or zero edits) away.
func oneAway(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			// add
			sa1 := s1[:i] + string(s2[i]) + s1[i:]
			sa2 := s2[:i] + string(s1[i]) + s2[i:]
			if sa1 == s2 || sa2 == s1 {
				return true
			}

			// remove
			sd1 := s1[:i] + s1[i+1:]
			sd2 := s2[:i] + s2[i+1:]
			if sd1 == s2 || sd2 == s1 {
				return true
			}

			// replace
			sr1 := s1[:i] + string(s2[i]) + s1[i+1:]
			sr2 := s2[:i] + string(s1[i]) + s2[i+1:]
			if sr1 == s2 || sr2 == s1 {
				return true
			}

			return false
		}
	}
	return true
}

// 1.6 String Compression
// Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3.
// If the "compressed" string would not become smaller than the original string,
// your method should return the original string.
// You can assume the string has only uppercase and lowercase letters (a - z).
func compressString(s string) string {
	var r strings.Builder
	r.WriteByte(s[0])
	count := 1
	for i := 1; i < len(s); i++ {
		if r.Len() >= len(s) {
			return s
		}

		if s[i] == s[i-1] {
			count++
		} else {
			r.WriteString(strconv.Itoa(count))
			r.WriteByte(s[i])
			count = 1
		}
	}

	if count != 0 {
		r.WriteString(strconv.Itoa(count))
	}

	return r.String()
}

// 1.7 Rotate Matrix
// Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
func rotateMatrix(m [][]int) [][]int {
	res := make([][]int, len(m))
	for i := range res {
		res[i] = make([]int, len(m))
	}

	for y := range m {
		rx := len(m) - 1 - y
		for x := range m {
			res[x][rx] = m[y][x]
		}
	}

	return res
}

// Extra
func reverse(str string) string {
	l, b := len(str), strings.Builder{}
	b.Grow(l)
	for i := l - 1; i >= 0; i-- {
		b.WriteByte(str[i])
	}

	return b.String()
}

func mergeSortedArrays(a, b []int) []int {
	m := []int{}
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i >= len(a) {
			m = append(m, b[j:]...)
			break
		} else if j >= len(b) {
			m = append(m, a[i:]...)
			break
		}

		if a[i] <= b[j] {
			m = append(m, a[i])
			i++
		} else if b[j] < a[i] {
			m = append(m, b[j])
			j++
		}
	}

	return m
}
