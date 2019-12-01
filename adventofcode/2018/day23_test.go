package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay23Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"pos=<0,0,0>, r=4",
				"pos=<1,0,0>, r=1",
				"pos=<4,0,0>, r=3",
				"pos=<0,2,0>, r=1",
				"pos=<0,5,0>, r=3",
				"pos=<0,0,3>, r=1",
				"pos=<1,1,1>, r=1",
				"pos=<1,1,2>, r=1",
				"pos=<1,3,1>, r=1",
			},
			7,
		},
	}

	for _, c := range cases {
		got := day23Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay23NewNanobot(t *testing.T) {
	cases := map[string]nanobot{
		"pos=<0,0,0>, r=4": nanobot{point3d{0, 0, 0}, 4},
		"pos=<1,0,0>, r=1": nanobot{point3d{1, 0, 0}, 1},
		"pos=<4,0,0>, r=3": nanobot{point3d{4, 0, 0}, 3},
		"pos=<0,2,0>, r=1": nanobot{point3d{0, 2, 0}, 1},
		"pos=<0,5,0>, r=3": nanobot{point3d{0, 5, 0}, 3},
		"pos=<0,0,3>, r=1": nanobot{point3d{0, 0, 3}, 1},
		"pos=<1,1,1>, r=1": nanobot{point3d{1, 1, 1}, 1},
		"pos=<1,1,2>, r=1": nanobot{point3d{1, 1, 2}, 1},
		"pos=<1,3,1>, r=1": nanobot{point3d{1, 3, 1}, 1},
	}

	for k, v := range cases {
		got := newNanobot(k)
		if !reflect.DeepEqual(got, v) {
			t.Errorf("case %s got %v, want %v", k, got, v)
		}
	}
}

func TestDay23Part1Final(t *testing.T) {
	in, err := readStrings("./day23_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day23Part1(*in)
	want := 219
	fmt.Printf("Day 23, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay23Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"pos=<10,12,12>, r=2",
				"pos=<12,14,12>, r=2",
				"pos=<16,12,12>, r=4",
				"pos=<14,14,14>, r=6",
				"pos=<50,50,50>, r=200",
				"pos=<10,10,10>, r=5",
			},
			36,
		},
	}

	for _, c := range cases {
		got := day23Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay23RangeIntersects(t *testing.T) {
	cases := []struct {
		b1   nanobot
		b2   nanobot
		want bool
	}{
		{
			nanobot{point3d{10, 12, 12}, 2},
			nanobot{point3d{10, 14, 12}, 2},
			true,
		},
		{
			nanobot{point3d{10, 12, 12}, 5},
			nanobot{point3d{11, 12, 12}, 2},
			true,
		},
		{
			nanobot{point3d{10, 12, 12}, 2},
			nanobot{point3d{10, 17, 12}, 2},
			false,
		},
		{
			nanobot{point3d{10, 12, 12}, 1},
			nanobot{point3d{10, 12, 15}, 1},
			false,
		},
	}

	for _, c := range cases {
		got := c.b1.rangeIntersects(c.b2)
		if got != c.want {
			t.Errorf("%v, %v, got %v, want %v", c.b1, c.b2, got, c.want)
		}
	}
}

func TestDay23Part2Final(t *testing.T) {
	in, err := readStrings("./day23_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day23Part2(*in)
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 23, part 2 answer: %v\n", got)
}
