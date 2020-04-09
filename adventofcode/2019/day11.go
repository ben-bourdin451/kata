package adventofcode

import "fmt"

const (
	black = 0
	white = 1
)

type hullbot struct {
	pos       point
	b         bounds
	direction int
	painted   map[point]int
}

func newHullbot() hullbot {
	return hullbot{point{0, 0}, bounds{0, 0, 0, 0}, up, make(map[point]int)}
}

func (b *hullbot) turnLeft() {
	b.direction--
	if b.direction < 0 {
		b.direction = 3
	}
}

func (b *hullbot) turnRight() {
	b.direction++
	if b.direction > 3 {
		b.direction = 0
	}
}

func (b *hullbot) getColor() int {
	if c, ok := b.painted[b.pos]; ok {
		return c
	}

	return black
}

func (b *hullbot) paint() {
	for y := b.b.minY; y <= b.b.maxY; y++ {
		for x := b.b.minX; x <= b.b.maxX; x++ {
			if color, ok := b.painted[point{x, y}]; ok && color == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (b *hullbot) start(mem []int64, in chan int64, done chan bool) {
	out := make(chan int64)

	go intcode(mem, in, out)
	for color := range out { // wait until intcode halts (out chan is closed)
		b.painted[b.pos] = int(color)
		newDir := <-out
		if newDir == 0 {
			b.turnLeft()
		} else {
			b.turnRight()
		}
		b.pos = b.pos.nextPoint(b.direction, 1)
		b.b.update(b.pos)
		in <- int64(b.getColor())
	}

	done <- true
}

func day11Part1(argv string) int {
	mem := initCodes(argv)

	bot := newHullbot()
	in, done := make(chan int64, 1), make(chan bool, 1)
	in <- 0 // send black as 1st input
	bot.start(mem, in, done)

	<-done // wait for bot to be finished

	return len(bot.painted)
}

func day11Part2(argv string) {
	mem := initCodes(argv)

	bot := newHullbot()
	in, done := make(chan int64, 1), make(chan bool, 1)
	in <- 1 // send white as 1st input
	bot.start(mem, in, done)

	<-done // wait for bot to be finished

	bot.paint()
}
