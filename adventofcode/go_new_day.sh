if [ $# -eq 0 ]; then
  echo "What day?\n"
	exit 0
fi

ifile="day${1}_input.txt"
tfile="day${1}_test.go"
cfile="day${1}.go"

touch $ifile

cat <<EOT >> $cfile
package adventofcode

func day$1Part1(in []string) int {
	return 0
}

func day$1Part2(in []string) int {
	return 0
}
EOT

cat <<EOT >> $tfile
package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay$1Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"",
			},
			0,
		},
	}

	for _, c := range cases {
		got := day$1Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay$1Part1Final(t *testing.T) {
	in, err := readStrings("./day$1_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day$1Part1(in)
	want := 0
	fmt.Printf("Day $1, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay$1Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"",
			},
			0,
		},
	}

	for _, c := range cases {
		got := day$1Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay$1Part2Final(t *testing.T) {
	in, err := readStrings("./day$1_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day$1Part2(in)
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day $1, part 2 answer: %v\n", got)
}
EOT