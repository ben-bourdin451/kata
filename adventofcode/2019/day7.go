package adventofcode

import (
	"reflect"
	"sync"
)

type amp struct {
	phase   int
	in, out chan int
}

func (a *amp) start(mem []int, wg *sync.WaitGroup) {
	go func() {
		intcode(mem, a.in, a.out)
		wg.Done()
	}()
	a.in <- a.phase
}

func amplifyThrusters(memory, phases []int) int {
	signal := 0
	for _, perm := range permutations(phases, []int{}) {
		in := make(chan int, 1)
		amps := []amp{}
		var wg sync.WaitGroup
		for i, phase := range perm {
			if i == 0 {
				amps = append(amps, amp{phase, in, make(chan int)})
			} else if i == 4 {
				amps = append(amps, amp{phase, amps[i-1].out, in})
			} else {
				amps = append(amps, amp{phase, amps[i-1].out, make(chan int)})
			}
			cpMem := make([]int, len(memory))
			reflect.Copy(reflect.ValueOf(cpMem), reflect.ValueOf(memory))
			wg.Add(1)
			amps[i].start(cpMem, &wg)
		}
		in <- 0
		wg.Wait()
		if output := <-in; output > signal {
			signal = output
		}
	}

	return signal
}

func day7Part1(argv string) int {
	memory := initCodes(argv)

	return amplifyThrusters(memory, []int{0, 1, 2, 3, 4})
}

func day7Part2(in string) int {
	memory := initCodes(in)

	return amplifyThrusters(memory, []int{5, 6, 7, 8, 9})
}
