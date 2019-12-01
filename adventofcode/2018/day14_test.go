package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	cases := []struct {
		in   int
		want []int
	}{
		{5, []int{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}},
		{9, []int{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}},
		{18, []int{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}},
		{2018, []int{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}},
	}

	for _, c := range cases {
		got := day14Part1(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay14Part1Final(t *testing.T) {
	got := day14Part1(360781)
	want := []int{6, 5, 2, 1, 5, 7, 1, 0, 1, 0}
	fmt.Printf("Day 14, part 1 answer: %v\n", got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay14Part2(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 1, 2, 4, 5, 1}, 5},
		{[]int{5, 1, 5, 8, 9, 1}, 9},
		{[]int{9, 2, 5, 1, 0, 7}, 18},
		{[]int{5, 9, 4, 1, 4, 2}, 2018},
	}

	for _, c := range cases {
		got := day14Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay14Part2Final(t *testing.T) {
	got := day14Part2([]int{3, 6, 0, 7, 8, 1})
	want := 20262967
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 14, part 2 answer: %v\n", got)
}
