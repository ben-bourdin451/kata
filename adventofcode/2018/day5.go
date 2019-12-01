package adventofcode

import (
	"math"
	"strings"
)

func day5Part1(in []string) int {
	reduced := true
	for reduced {
		in, reduced = pReduce(in)
	}
	return len(in)
}

func pReduce(s []string) ([]string, bool) {
	cL, cR := make(chan []string), make(chan []string)
	l, r := s[:len(s)/2], s[len(s)/2:]
	lenL, lenR := len(l), len(r)

	go rReduce(cL, l)
	go rReduce(cR, r)

	l, r = <-cL, <-cR

	replaced := false
	if lenL > len(l) || lenR > len(r) {
		replaced = true
	}

	if len(l) == 0 || len(r) == 0 {
		return append(l, r...), replaced
	}

	if areOpposite(l[len(l)-1], r[0]) {
		s = append(l[:len(l)-1], r[1:]...)
		replaced = true
	} else {
		s = append(l, r...)
	}

	return s, replaced
}

func rReduce(c chan []string, s []string) {
	if len(s) >= 20 {
		res, _ := pReduce(s)
		c <- res
	} else {
		c <- reduce(s)
	}
}

func reduce(s []string) []string {
	if len(s) < 2 {
		return s

	} else if areOpposite(s[0], s[1]) {
		if len(s) == 2 {
			return []string{}
		}
		return reduce(s[2:])

	} else {
		return append(s[:1], reduce(s[1:])...)
	}
}

func areOpposite(a, b string) bool {
	return math.Abs(float64(a[0])-float64(b[0])) == 32
}

func day5Part2(input string) int {
	ch := make(chan int)
	for i := 65; i <= 90; i++ {
		go func(in string, n int) {
			s := strings.Replace(in, string(n), "", -1)
			s = strings.Replace(s, string(n+32), "", -1)
			r := day5Part1(strings.Split(s, ""))
			ch <- r
		}(input, i)
	}

	best := len(input)
	for i := 65; i <= 90; i++ {
		r := <-ch
		if r < best {
			best = r
		}
	}
	return best
}
