package adventofcode

import (
	"reflect"
	"strconv"
	"strings"
)

func day14Part1(r int) []int {
	recipes := []int{3, 7}
	e1, e2 := 0, 1
	lenR := len(recipes)
	for lenR < r+10 {
		s := strconv.Itoa(recipes[e1] + recipes[e2])
		for _, n := range strings.Split(s, "") {
			num, _ := strconv.Atoi(n)
			recipes = append(recipes, num)
		}
		lenR = len(recipes)

		e1 += recipes[e1] + 1
		e1 = e1 % lenR

		e2 += recipes[e2] + 1
		e2 = e2 % lenR
	}

	return recipes[r : r+10]
}

func day14Part2(in []int) int {
	recipes := []int{3, 7}
	e1, e2 := 0, 1
	lenR := len(recipes)
	for {
		s := strconv.Itoa(recipes[e1] + recipes[e2])
		for _, n := range strings.Split(s, "") {
			num, _ := strconv.Atoi(n)
			recipes = append(recipes, num)
			lenR = len(recipes)

			if lenR > 6 && reflect.DeepEqual(in, recipes[lenR-6:lenR]) {
				return lenR - 6
			}
		}

		e1 += recipes[e1] + 1
		e1 = e1 % lenR

		e2 += recipes[e2] + 1
		e2 = e2 % lenR
	}
}
