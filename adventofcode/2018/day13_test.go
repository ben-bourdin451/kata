package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want *point
	}{
		{
			[]string{
				`/->-\        `,
				`|   |  /----\`,
				`| /-+--+-\  |`,
				`| | |  | v  |`,
				`\-+-/  \-+--/`,
				`  \------/   `,
			},
			&point{7, 3},
		},
	}

	for _, c := range cases {
		got := day13Part1(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay13Part1Final(t *testing.T) {
	in, err := readStrings("./day13_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day13Part1(*in)
	want := &point{103, 85}
	fmt.Printf("Day 13, part 1 answer: %v\n", got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay13Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want *point
	}{
		{
			[]string{
				`/>-<\  `,
				`|   |  `,
				`| /<+-\`,
				`| | | v`,
				`\>+</ |`,
				`  |   ^`,
				`  \<->/`,
			},
			&point{6, 4},
		},
	}

	for _, c := range cases {
		got := day13Part2(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay13Part2Final(t *testing.T) {
	in, err := readStrings("./day13_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day13Part2(*in)
	want := &point{88, 64}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 13, part 2 answer: %v\n", got)
}
