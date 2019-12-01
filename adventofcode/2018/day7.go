package adventofcode

import (
	"regexp"
	"sort"
	"strings"
)

type task struct {
	id         string
	deps       []*task
	calls      []*task
	isDone     bool
	inProgress bool
}

func (n *task) findNextCalls() []*task {
	next := []*task{}
	for _, call := range n.calls {
		if call.areDepsDone() && !call.inProgress && !call.isDone {
			next = append(next, call)
		}
	}

	return next
}

func (n *task) areDepsDone() bool {
	for _, dep := range n.deps {
		if !dep.isDone {
			return false
		}
	}
	return true
}

func day7Part1(input []string) string {
	taskMap := parseInput(input)

	r := []string{}
	callStack := findNodesWithNoDeps(taskMap)
	for len(callStack) > 0 {
		task := callStack[0]
		task.isDone = true
		r = append(r, task.id)
		callStack = append(callStack[1:], task.findNextCalls()...)
		sortNodes(callStack)
	}

	return strings.Join(r, "")
}

func parseInput(input []string) map[string]*task {
	taskMap := make(map[string]*task, len(input))
	for _, line := range input {
		s1, s2 := parseLine(line)

		if _, ok := taskMap[s1]; !ok {
			taskMap[s1] = &task{
				id:         s1,
				deps:       []*task{},
				calls:      []*task{},
				isDone:     false,
				inProgress: false,
			}
		}

		if _, ok := taskMap[s2]; !ok {
			taskMap[s2] = &task{
				id:         s2,
				deps:       []*task{},
				calls:      []*task{},
				isDone:     false,
				inProgress: false,
			}
		}

		taskMap[s1].calls = append(taskMap[s1].calls, taskMap[s2])
		taskMap[s2].deps = append(taskMap[s2].deps, taskMap[s1])
	}

	return taskMap
}

func parseLine(l string) (string, string) {
	r, _ := regexp.Compile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)
	return r.FindStringSubmatch(l)[1], r.FindStringSubmatch(l)[2]
}

func sortNodes(tasks []*task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].id < tasks[j].id
	})
}

func findNodesWithNoDeps(taskMap map[string]*task) []*task {
	n := []*task{}
	for _, v := range taskMap {
		if len(v.deps) == 0 {
			n = append(n, v)
		}
	}
	sortNodes(n)
	return n
}

func day7Part2(input []string, numElves, stepCost int) int {
	taskMap := parseInput(input)

	elves := []*worker{}
	for i := 0; i < numElves; i++ {
		elves = append(elves, &worker{
			id:   i,
			cost: stepCost,
			busy: false,
		})
	}

	s := 0
	callStack := findNodesWithNoDeps(taskMap)
	for busy := false; len(callStack) > 0 || busy; {
		busy = false
		// fmt.Printf("\nS%v ", s)
		for _, elf := range elves {
			if !elf.busy && len(callStack) > 0 {
				callStack = elf.take(callStack)
			}

			elf.work()
		}

		for _, elf := range elves {
			if elf.slot != nil && !elf.busy {
				callStack = append(callStack, elf.slot.findNextCalls()...)
				elf.slot = nil
			}

			busy = busy || elf.busy
		}

		sortNodes(callStack)
		s++
	}

	return s
}

type worker struct {
	id                 int
	slot               *task
	count, limit, cost int
	busy               bool
}

func (w *worker) take(work []*task) []*task {
	w.start(work[0])
	return work[1:]
}

func (w *worker) start(n *task) {
	w.limit = int(n.id[0]) - 64 + w.cost // ASCII values 65-90 --> A=1second
	w.count = 0
	w.busy = true

	n.inProgress = true
	w.slot = n
}

func (w *worker) work() {
	if w.limit <= 0 {
		// fmt.Printf("\t#%v\t.", w.id)
		return
	}

	w.count++
	// fmt.Printf("\t#%v\t%v:%v", w.id, w.slot.id, w.count)

	if w.count >= w.limit {
		w.slot.isDone = true
		w.slot.inProgress = false
		w.busy = false
		w.limit = 0
		w.count = 0
	}
}
