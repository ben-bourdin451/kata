package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"123456789012", 1},
	}

	for _, c := range cases {
		got := day8Part1(c.in, 3, 2)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay8Part1Final(t *testing.T) {
	in, err := readStrings("./day8_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day8Part1(in[0], 25, 6)
	fmt.Printf("Day 8, part 1 answer: %v\n", got)
	want := 2016
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay8Part2Final(t *testing.T) {
	in, err := readStrings("./day8_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	fmt.Printf("Day 8, part 2 answer:\n")
	day8Part2(in[0], 25, 6)
}
