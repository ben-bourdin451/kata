package adventofcode

import "testing"

func Test_bounds_contains(t *testing.T) {
	type fields struct {
		minX int
		minY int
		maxX int
		maxY int
	}
	type args struct {
		p point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"0 bounds",
			fields{0, 0, 0, 0},
			args{point{0, 0}},
			false,
		},
		{
			"inside bounds",
			fields{0, 0, 10, 10},
			args{point{5, 2}},
			true,
		},
		{
			"outside bounds",
			fields{0, 0, 10, 10},
			args{point{11, 2}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bounds{
				minX: tt.fields.minX,
				minY: tt.fields.minY,
				maxX: tt.fields.maxX,
				maxY: tt.fields.maxY,
			}
			if got := b.contains(tt.args.p); got != tt.want {
				t.Errorf("bounds.contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
