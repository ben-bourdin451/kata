package cci

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
