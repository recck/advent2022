package puzzles

import (
	"bufio"
	"image"
	"os"
	"strconv"
	"strings"
)

var allAdjacentPoints = []image.Point{
	{-1, 1},  // top left
	{0, 1},   // top
	{1, 1},   // top right
	{-1, 0},  // left
	{0, 0},   // same
	{1, 0},   // right
	{-1, -1}, // bottom left
	{0, -1},  // bottom
	{1, -1},  // bottom right
}

type movement struct {
	direction image.Point
	numSteps  int
}

type Day9 struct {
	head, tail   image.Point
	tailSeen     map[image.Point]int
	instructions []movement
}

func (d *Day9) isTailBehind() bool {
	for _, ap := range allAdjacentPoints {
		if d.head.Add(ap).Eq(d.tail) {
			return true
		}
	}

	return false
}

func (d *Day9) addTailSeen(t image.Point) {
	if _, ok := d.tailSeen[t]; !ok {
		d.tailSeen[t] = 1
	}
}

func (d *Day9) Part1() any {
	for i := range d.instructions {
		for j := 0; j < d.instructions[i].numSteps; j++ {
			d.head = d.head.Add(d.instructions[i].direction)
			if !d.isTailBehind() {
				d.tail = d.head.Sub(d.instructions[i].direction)
				d.addTailSeen(d.tail)
			}
		}
	}

	return len(d.tailSeen)
}

func (d *Day9) Part2() any {
	return 0
}

func (d *Day9) LoadInput(file *os.File) error {
	var instructions []movement
	d.tailSeen = map[image.Point]int{{0, 0}: 1}
	d.head = image.Point{}
	d.tail = image.Point{}

	instMap := map[string]image.Point{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Fields(line)
		steps, _ := strconv.Atoi(chunks[1])
		instructions = append(instructions, movement{
			direction: instMap[chunks[0]],
			numSteps:  steps,
		})
	}

	d.instructions = instructions

	return nil
}
