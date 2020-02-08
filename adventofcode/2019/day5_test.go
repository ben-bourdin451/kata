package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay5Part1Final(t *testing.T) {
	in, err := readStrings("./day5_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day5(in[0], 1)
	want := int64(15508323)
	fmt.Printf("Day 5, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay5Part2Final(t *testing.T) {
	in, err := readStrings("./day5_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day5(in[0], 5)
	want := int64(9006327)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 5, part 2 answer: %v\n", got)
}
