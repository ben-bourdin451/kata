package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay18Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				".#.#...|#.",
				".....#|##|",
				".|..|...#.",
				"..|#.....#",
				"#.#|||#|#|",
				"...#.||...",
				".|....|...",
				"||...#|.#|",
				"|.||||..|.",
				"...#.|..|.",
			},
			1147,
		},
	}

	for _, c := range cases {
		got := day18(c.in, 10)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay18Part1Final(t *testing.T) {
	in, err := readStrings("./day18_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day18(*in, 10)
	want := 531417
	fmt.Printf("Day 18, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay18Part2Final(t *testing.T) {
	in, err := readStrings("./day18_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day18(*in, 1000000000)
	want := 205296
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 18, part 2 answer: %v\n", got)
}
