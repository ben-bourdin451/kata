package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	cases := []struct {
		in   []*sample
		want int
	}{
		{
			[]*sample{&sample{
				[]int{3, 2, 1, 1},
				[]int{9, 2, 1, 2},
				[]int{3, 2, 2, 1},
			}},
			1,
		},
	}

	for _, c := range cases {
		got := day16Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay16BehaviourCount(t *testing.T) {
	cases := []struct {
		in   *sample
		want int
	}{
		{
			&sample{
				[]int{3, 2, 1, 1},
				[]int{9, 2, 1, 2},
				[]int{3, 2, 2, 1},
			},
			3,
		},
	}

	for _, c := range cases {
		got, _ := behaviourCount(c.in, initBehaviours())
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay16ParseSamples(t *testing.T) {
	cases := []struct {
		in   []string
		want *sample
	}{
		{
			[]string{
				"",
				"Before: [3, 2, 1, 1]",
				"9 2 1 2",
				"After:  [3, 2, 2, 1]",
				"",
			},
			&sample{
				[]int{3, 2, 1, 1},
				[]int{9, 2, 1, 2},
				[]int{3, 2, 2, 1},
			},
		},
		{
			[]string{
				"",
				"Before: [0, 1, 2, 3]",
				"13 2 1 2",
				"After:  [3, 2, 1, 0]",
			},
			&sample{
				[]int{0, 1, 2, 3},
				[]int{13, 2, 1, 2},
				[]int{3, 2, 1, 0},
			},
		},
	}

	for _, c := range cases {
		got := parseSamples(c.in)

		if !reflect.DeepEqual(got[0], c.want) {
			t.Errorf("got %v, want %v", got[0], c.want)
		}
	}
}

func TestDay16ParseRegister(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{"1 2 1 0", []int{1, 2, 1, 0}},
		{"9 1 0 2", []int{9, 1, 0, 2}},
		{"10 2 3 2", []int{10, 2, 3, 2}},
		{"1 1 2 3", []int{1, 1, 2, 3}},
		{"15 0 3 3", []int{15, 0, 3, 3}},
	}

	for _, c := range cases {
		got := parseRegister(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay16Behaviours(t *testing.T) {
	cases := []struct {
		r    []int
		in   []int
		op   opFunc
		want []int
	}{
		{
			[]int{3, 2, 1, 1},
			[]int{9, 2, 1, 2},
			addr,
			[]int{3, 2, 3, 1},
		},
		{
			[]int{3, 2, 1, 1},
			[]int{9, 2, 1, 2},
			addi,
			[]int{3, 2, 2, 1},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 1, 2},
			mulr,
			[]int{3, 5, 10, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 2, 2},
			muli,
			[]int{3, 5, 4, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 1, 2},
			banr,
			[]int{3, 5, 2 & 5, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 2, 2},
			bani,
			[]int{3, 5, 2 & 2, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 1, 2},
			borr,
			[]int{3, 5, 2 | 5, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 2, 2},
			bori,
			[]int{3, 5, 2 | 2, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 1, 2},
			setr,
			[]int{3, 5, 2, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 6, 2, 2},
			seti,
			[]int{3, 5, 6, 4},
		},
		{
			[]int{3, 1, 2, 4},
			[]int{9, 2, 1, 2},
			gtir,
			[]int{3, 1, 1, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 1, 2},
			gtir,
			[]int{3, 5, 0, 4},
		},
		{
			[]int{3, 1, 2, 4},
			[]int{9, 2, 1, 2},
			gtri,
			[]int{3, 1, 1, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 3, 2},
			gtri,
			[]int{3, 5, 0, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 3, 2, 2},
			gtrr,
			[]int{3, 5, 1, 4},
		},
		{
			[]int{3, 5, 2, 1},
			[]int{9, 3, 2, 2},
			gtrr,
			[]int{3, 5, 0, 1},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 1, 2},
			eqir,
			[]int{3, 5, 0, 4},
		},
		{
			[]int{3, 1, 2, 4},
			[]int{9, 4, 3, 2},
			eqir,
			[]int{3, 1, 1, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 1, 2, 2},
			eqri,
			[]int{3, 5, 0, 4},
		},
		{
			[]int{3, 5, 2, 4},
			[]int{9, 2, 2, 2},
			eqri,
			[]int{3, 5, 1, 4},
		},
		{
			[]int{3, 5, 3, 4},
			[]int{9, 0, 2, 2},
			eqrr,
			[]int{3, 5, 1, 4},
		},
		{
			[]int{3, 5, 3, 4},
			[]int{9, 0, 1, 2},
			eqrr,
			[]int{3, 5, 0, 4},
		},
	}

	for _, c := range cases {
		original := []int{0, 0, 0, 0}
		copy(original[:], c.r)

		got := c.op(c.r, c.in[1:])

		if !reflect.DeepEqual(got, c.want) || !reflect.DeepEqual(c.r, original) {
			t.Errorf("got %v, want %v", got, c.want)
			t.Errorf("had %v, got %v", original, c.r)
		}
	}
}

func TestDay16Part1Final(t *testing.T) {
	in, err := readStrings("./day16_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	s := parseSamples(*in)

	got := day16Part1(s)
	want := 588
	fmt.Printf("Day 16, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay16Part2Final(t *testing.T) {
	samples, err := readStrings("./day16_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	in, err := readStrings("./day16_input_2.txt")
	if err != nil {
		t.Error("Error while reading input 2", err)
	}

	s := parseSamples(*samples)

	got := day16Part2(s, *in)
	want := 627
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 16, part 2 answer: %v\n", got)
}
