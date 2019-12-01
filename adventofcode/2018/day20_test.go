package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay20Part1(t *testing.T) {
	cases := map[string]int{
		"^WNE$":                                                             3,
		"^ENWWW(NEEE|SSE(EE|N))$":                                           10,
		"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$":                         18,
		"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$":               23,
		"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$": 31,
	}

	for k, v := range cases {
		got := day20Part1(k)
		if got != v {
			t.Errorf("%s\ngot %v, want %v", k, got, v)
		}
	}
}

func TestDay20Part1Final(t *testing.T) {
	in, err := readStringsPlain("./day20_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day20Part1(in[0])
	want := 3314
	fmt.Printf("Day 20, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay20Part2(t *testing.T) {
	cases := map[string]struct {
		distance int
		want     int
	}{
		"^WNE$": {
			2,
			2,
		},
		"^ENWWW(NEEE|SSE(EE|N))$": {
			8,
			6,
		},
		"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$": {
			9,
			15,
		},
		"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$": {
			21,
			4,
		},
	}

	for k, v := range cases {
		got := day20Part2(k, v.distance)
		if got != v.want {
			t.Errorf("Distance[%v], Regex[%s]\ngot %v, want %v", v.distance, k, got, v.want)
		}
	}
}

func TestDay20Part2Final(t *testing.T) {
	in, err := readStringsPlain("./day20_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day20Part2(in[0], 1000)
	want := 8550
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 20, part 2 answer: %v\n", got)
}
