package adventofcode

import (
	"reflect"
	"sync"
)

type amp struct {
	phase   int64
	in, out chan int64
}

func (a *amp) start(mem []int64, wg *sync.WaitGroup) {
	go func() {
		intcode(mem, a.in, a.out)
		wg.Done()
	}()
	a.in <- a.phase
}

func amplifyThrusters(memory []int64, phases []int) int64 {
	signal := int64(0)
	for _, perm := range permutations(phases, []int{}) {
		in := make(chan int64, 1)
		amps := []amp{}
		var wg sync.WaitGroup
		for i, phase := range perm {
			if i == 0 {
				amps = append(amps, amp{int64(phase), in, make(chan int64)})
			} else if i == 4 {
				amps = append(amps, amp{int64(phase), amps[i-1].out, in})
			} else {
				amps = append(amps, amp{int64(phase), amps[i-1].out, make(chan int64)})
			}
			cpMem := make([]int64, len(memory))
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

func day7Part1(argv string) int64 {
	memory := initCodes(argv)

	return amplifyThrusters(memory, []int{0, 1, 2, 3, 4})
}

func day7Part2(in string) int64 {
	memory := initCodes(in)

	return amplifyThrusters(memory, []int{5, 6, 7, 8, 9})
}
