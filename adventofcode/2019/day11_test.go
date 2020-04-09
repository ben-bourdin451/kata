package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay11Part1Final(t *testing.T) {
	in, err := readStrings("./day11_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day11Part1(in[0])
	want := 1932
	fmt.Printf("Day 11, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay11Part2Final(t *testing.T) {
	in, err := readStrings("./day11_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	fmt.Println("Day 11, part 2 answer:")
	day11Part2(in[0])
}
