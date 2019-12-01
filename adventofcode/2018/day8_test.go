package adventofcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{
			"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2",
			138,
		},
	}

	for _, c := range cases {
		got := day8Part1(strings.Split(c.in, " "))
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

	got := day8Part1(strings.Split((*in)[0], " "))
	want := 42501
	fmt.Printf("Day 8, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay8Part2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{
			"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2",
			66,
		},
	}

	for _, c := range cases {
		got := day8Part2(strings.Split(c.in, " "))
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay8Part2Final(t *testing.T) {
	in, err := readStrings("./day8_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day8Part2(strings.Split((*in)[0], " "))
	want := 30857
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 8, part 2 answer: %v\n", got)
}
