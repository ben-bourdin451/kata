package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay9Part1(t *testing.T) {
	cases := []struct {
		players    int
		lastMarble int
		highScore  int
	}{
		{9, 25, 32},
		// {9, 50, 32},
		{10, 1618, 8317},   // 10 players; last marble is worth 1618 points: high score is 8317
		{13, 7999, 146373}, // 13 players; last marble is worth 7999 points: high score is 146373
		{17, 1104, 2764},   // 17 players; last marble is worth 1104 points: high score is 2764
		{21, 6111, 54718},  // 21 players; last marble is worth 6111 points: high score is 54718
		{30, 5807, 37305},  // 30 players; last marble is worth 5807 points: high score is 37305
	}

	for _, c := range cases {
		got := day9Part1(c.players, c.lastMarble)
		if got != c.highScore {
			t.Errorf("got %v, want %v", got, c.highScore)
		}
	}
}

func TestDay9Part1Final(t *testing.T) {
	got := day9Part1(416, 71617) // 416 players; last marble is worth 71617 points
	want := 436720
	fmt.Printf("Day 9, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay9Part2Final(t *testing.T) {
	got := day9Part1(416, 7161700) // 416 players; last marble is worth 71617 points
	want := 3527845091
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 9, part 2 answer: %v\n", got)
}
