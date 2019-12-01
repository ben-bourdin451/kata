package adventofcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func readIntegers(path string) (*[]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var in []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		in = append(in, n)
	}
	return &in, nil
}

func TestDay1Part1(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}
	for _, c := range cases {
		got := day1Part1(&c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay1Part1Final(t *testing.T) {
	in, err := readIntegers("./day1_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day1Part1(in)
	want := 411
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 1, part 1 answer: %v\n", got)
}

func TestDay1Ex2(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}
	for _, c := range cases {
		got := day1Part2(&c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay1Part2Final(t *testing.T) {
	in, err := readIntegers("./day1_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day1Part2(in)
	want := 56360
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 1, part 2 answer: %v\n", got)
}
