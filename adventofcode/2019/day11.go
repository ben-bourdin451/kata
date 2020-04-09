package adventofcode

const (
	black = 0
	white = 1
)

type hullbot struct {
	pos       point
	direction int
	painted   map[point]int
}

func newHullbot() hullbot {
	return hullbot{point{0, 0}, up, make(map[point]int)}
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

// func (b *hullbot) paint(c int) {
// 	b.painted[b.pos] = c
// }

func (b *hullbot) start(mem []int64, done chan bool) {
	in, out := make(chan int64, 1), make(chan int64)

	go intcode(mem, in, out)
	in <- 0                  // send black as 1st input
	for color := range out { // wait until intcode halts (out chan is closed)
		b.painted[b.pos] = int(color)
		newDir := <-out
		if newDir == 0 {
			b.turnLeft()
		} else {
			b.turnRight()
		}
		b.pos = b.pos.nextPoint(b.direction, 1)
		in <- int64(b.getColor())
	}

	done <- true
}

func day11Part1(argv string) int {
	mem := initCodes(argv)

	bot := newHullbot()
	done := make(chan bool, 1)
	bot.start(mem, done)

	<-done // wait for bot to be finished

	return len(bot.painted)
}

func day11Part2(in []string) int {
	return 0
}
