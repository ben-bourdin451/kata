package adventofcode

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type sleepRecord struct {
	total   int
	minutes []int // [0,12,3,4,...]{60}
}

func day4Part1(in []string) int {
	sortInput(in)

	var guard int
	var prevTime time.Time
	guardRecords := make(map[int]*sleepRecord, 2500)
	for _, text := range in {
		// Parse time
		time, _ := getTime(text)

		if hasGuardID(text) {
			guard, _ = getGuardID(text)
			if _, ok := guardRecords[guard]; !ok {
				guardRecords[guard] = &sleepRecord{0, make([]int, 60)}
			}
		} else if strings.Contains(text, "wakes") {
			for i := prevTime.Minute(); i < time.Minute(); i++ {
				guardRecords[guard].minutes[i]++
				guardRecords[guard].total++
			}
		}
		prevTime = time
	}

	id := findGuardMostSleep(guardRecords)
	minute, _ := findMostSleptMinute(guardRecords[id])
	return id * minute
}

func sortInput(in []string) {
	sort.Slice(in, func(i, j int) bool {
		iTime, _ := getTime(in[i])
		jTime, _ := getTime(in[j])
		return iTime.Before(jTime)
	})
}

func getTime(text string) (time.Time, error) {
	format := "[2006-01-02 15:04]"
	r := regexp.MustCompile(`^\[(?P<time>.+)\]`)
	t := r.FindStringSubmatch(text)[0]
	return time.Parse(format, t)
}

func hasGuardID(text string) bool {
	rGuard := regexp.MustCompile(`Guard #(?P<id>\d+) begins shift`)
	return rGuard.MatchString(text)
}

func getGuardID(text string) (int, error) {
	rGuard := regexp.MustCompile(`Guard #(?P<id>\d+) begins shift`)
	return strconv.Atoi(rGuard.FindStringSubmatch(text)[1])
}

func findGuardMostSleep(records map[int]*sleepRecord) int {
	best, id := 0, 0
	for k, v := range records {
		if v.total > best {
			best = v.total
			id = k
		}
	}
	return id
}

func findMostSleptMinute(record *sleepRecord) (int, int) {
	sleepy, minute := 0, 0
	for i, v := range record.minutes {
		if v > sleepy {
			sleepy = v
			minute = i
		}
	}
	return minute, sleepy
}

func day4Part2(in []string) int {
	sortInput(in)

	var guard int
	var prevTime time.Time
	guardRecords := make(map[int]*sleepRecord, 2500)
	for _, text := range in {
		// Parse time
		time, _ := getTime(text)

		if hasGuardID(text) {
			guard, _ = getGuardID(text)
			if _, ok := guardRecords[guard]; !ok {
				guardRecords[guard] = &sleepRecord{0, make([]int, 60)}
			}
		} else if strings.Contains(text, "wakes") {
			for i := prevTime.Minute(); i < time.Minute(); i++ {
				guardRecords[guard].minutes[i]++
				guardRecords[guard].total++
			}
		}
		prevTime = time
	}

	best, id, minute := 0, 0, 0
	for guard, record := range guardRecords {
		m, v := findMostSleptMinute(record)

		if v > best {
			best = v
			minute = m
			id = guard
		}
	}

	return id * minute
}
