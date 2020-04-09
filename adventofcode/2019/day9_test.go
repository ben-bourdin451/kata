package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay9Part1(t *testing.T) {
	cases := []struct {
		in   string
		want []int64
	}{
		{
			"104,1125899906842624,99",
			[]int64{1125899906842624},
		},
		{
			"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
			[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			"1102,34915192,34915192,7,4,7,99,0",
			[]int64{1219070632396864},
		},
	}

	for _, c := range cases {
		got := day9Part1(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay9Part1Final(t *testing.T) {
	in, err := readStrings("./day9_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day9Part1(in[0])
	want := []int64{2890527621}
	fmt.Printf("Day 9, part 1 answer: %v\n", got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay9Part2Final(t *testing.T) {
	in, err := readStrings("./day9_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day9Part2(in[0])
	want := []int64{66772}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 9, part 2 answer: %v\n", got)
}
