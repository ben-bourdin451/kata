package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay17Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"x=495, y=2..7",
				"y=7, x=495..501",
				"x=501, y=3..7",
				"x=498, y=2..4",
				"x=506, y=1..2",
				"x=498, y=10..13",
				"x=504, y=10..13",
				"y=13, x=498..504",
			},
			57,
		},
		{
			[]string{
				"x=492, y=3..12",
				"x=508, y=4..12",
				"y=13, x=492..508",
				"y=7, x=500..502",
				"y=10, x=500..502",
				"x=500, y=8..9",
				"x=502, y=8..9",
				"x=513, y=1..2",
			},
			152,
		},
		{
			[]string{
				"x=479, y=1..5",
				"x=490, y=4..7",
				"x=514, y=4..7",
				"y=7, x=491..513",
				"y=4, x=531..534",
				"x=532, y=11..25",
				"x=482, y=12..25",
				"y=25, x=483..531",
				"y=21, x=511..523",
				"y=15, x=511..523",
				"x=511, y=16..20",
				"x=523, y=16..20",
			},
			723,
		},
	}

	for _, c := range cases {
		got := day17Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay17Part1Final(t *testing.T) {
	in, err := readStrings("./day17_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day17Part1(*in)
	want := 27042
	fmt.Printf("Day 17, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay17Flow(t *testing.T) {
	clay := map[point]struct{}{
		point{495, 2}: {},
		point{495, 3}: {},
		point{495, 4}: {},
		point{495, 5}: {},
		point{495, 6}: {},
		point{495, 7}: {},
		point{496, 7}: {},
		point{497, 7}: {},
		point{498, 7}: {},
		point{498, 2}: {},
		point{498, 3}: {},
		point{498, 4}: {},
		point{499, 7}: {},
		point{500, 7}: {},
		point{501, 7}: {},
		point{501, 6}: {},
		point{501, 5}: {},
		point{501, 4}: {},
		point{501, 3}: {},
	}

	cases := []struct {
		s         *point
		d         int
		w         map[point]struct{}
		newSpring *point
	}{
		{&point{500, 6}, right, map[point]struct{}{}, nil},
		{&point{500, 6}, left, map[point]struct{}{}, nil},
		{&point{500, 3}, left, map[point]struct{}{point{499, 4}: {}, point{500, 4}: {}}, nil},
		{&point{500, 2}, right, map[point]struct{}{point{499, 3}: {}, point{500, 3}: {}}, &point{502, 2}},
	}

	for _, c := range cases {
		s := &soil{&bounds{510, 510, 400, 0}, clay, c.w, map[point]struct{}{}}
		spring := s.spread(c.s, c.d)
		if !reflect.DeepEqual(spring, c.newSpring) {
			t.Errorf("got %v, want %v", spring, c.newSpring)
		}
	}
}

func TestDay17ParseClay(t *testing.T) {
	cases := []struct {
		in   string
		want []*point
	}{
		{
			"x=495, y=2..7",
			[]*point{
				&point{495, 2},
				&point{495, 3},
				&point{495, 4},
				&point{495, 5},
				&point{495, 6},
				&point{495, 7},
			},
		},
		{
			"y=7, x=495..501",
			[]*point{
				&point{495, 7},
				&point{496, 7},
				&point{497, 7},
				&point{498, 7},
				&point{499, 7},
				&point{500, 7},
				&point{501, 7},
			},
		},
	}

	for _, c := range cases {
		got := parseClay(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay17Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"x=495, y=2..7",
				"y=7, x=495..501",
				"x=501, y=3..7",
				"x=498, y=2..4",
				"x=506, y=1..2",
				"x=498, y=10..13",
				"x=504, y=10..13",
				"y=13, x=498..504",
			},
			29,
		},
		{
			[]string{
				"x=492, y=3..12",
				"x=508, y=4..12",
				"y=13, x=492..508",
				"y=7, x=500..502",
				"y=10, x=500..502",
				"x=500, y=8..9",
				"x=502, y=8..9",
				"x=513, y=1..2",
			},
			123,
		},
		{
			[]string{
				"x=479, y=1..5",
				"x=490, y=4..7",
				"x=514, y=4..7",
				"y=7, x=491..513",
				"y=4, x=531..534",
				"x=532, y=11..25",
				"x=482, y=12..25",
				"y=25, x=483..531",
				"y=21, x=511..523",
				"y=15, x=511..523",
				"x=511, y=16..20",
				"x=523, y=16..20",
			},
			615,
		},
	}

	for _, c := range cases {
		got := day17Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay17Part2Final(t *testing.T) {
	in, err := readStrings("./day17_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day17Part2(*in)
	want := 22214
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 17, part 2 answer: %v\n", got)
}
