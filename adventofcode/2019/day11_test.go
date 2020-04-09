package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"",
			},
			0,
		},
	}

	for _, c := range cases {
		got := day11Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay11Part1Final(t *testing.T) {
	in, err := readStrings("./day11_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day11Part1(in)
	want := 0
	fmt.Printf("Day 11, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay11Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"",
			},
			0,
		},
	}

	for _, c := range cases {
		got := day11Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay11Part2Final(t *testing.T) {
	in, err := readStrings("./day11_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day11Part2(in)
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 11, part 2 answer: %v\n", got)
}
