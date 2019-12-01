package adventofcode

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func readStrings(path string) (*[]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var in []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return &in, nil
}

func readStringsPlain(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var in []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return in, nil
}

func TestDay2Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{
			"abcdef",
			"bababc",
			"abbcde",
			"abcccd",
			"aabcdd",
			"abcdee",
			"ababab",
		}, 12},
	}
	for _, c := range cases {
		got := day2Part1(&c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay2Part1Final(t *testing.T) {
	in, err := readStrings("./day2_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day2Part1(in)
	want := 7105
	fmt.Printf("Day 2, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay2Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{
			"abcde",
			"fghij",
			"klmno",
			"pqrst",
			"fguij",
			"axcye",
			"wvxyz",
		}, "fgij"},
	}
	for _, c := range cases {
		got := day2Part2(&c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay2Part2Final(t *testing.T) {
	in, err := readStrings("./day2_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day2Part2(in)
	want := "omlvgdokxfncvqyersasjziup"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 2, part 2 answer: %v\n", got)
}
