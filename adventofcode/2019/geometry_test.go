package adventofcode

import (
	"reflect"
	"testing"
)

func TestGeometryNewVector(t *testing.T) {
	origin := point{0, 0}

	cases := []struct {
		p1, p2 point
		want   vector
	}{
		{origin, point{0, 1}, vector{0, 1, 1}},
		{origin, point{0, 3}, vector{0, 1, 3}},
		{origin, point{4, 4}, vector{1, 1, 4}},
		{origin, point{2, 1}, vector{2, 1, 1}},
		{origin, point{4, 2}, vector{2, 1, 2}},
	}

	for _, c := range cases {
		got := newVector(c.p1, c.p2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
