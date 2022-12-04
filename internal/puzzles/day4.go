package puzzles

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	start, end int
}

func (a assignment) contains(a2 assignment) bool {
	return a2.start >= a.start && a2.end <= a.end
}

func (a assignment) overlaps(a2 assignment) bool {
	return a.start <= a2.end && a2.start <= a.end
}

func newAssignment(s string) assignment {
	bounds := strings.Split(s, "-")
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])

	return assignment{start: start, end: end}
}

type assignmentPair struct {
	left, right assignment
}

func (a assignmentPair) hasContainment() bool {
	return a.left.contains(a.right) || a.right.contains(a.left)
}

func (a assignmentPair) hasOverlap() bool {
	return a.left.overlaps(a.right) || a.right.overlaps(a.left)
}

type Day4 struct {
	assignmentPairs []assignmentPair
}

func (d *Day4) Part1() any {
	count := 0
	for _, ap := range d.assignmentPairs {
		if ap.hasContainment() {
			count++
		}
	}

	return count
}

func (d *Day4) Part2() any {
	count := 0
	for _, ap := range d.assignmentPairs {
		if ap.hasOverlap() {
			count++
		}
	}

	return count
}

func (d *Day4) LoadInput(file *os.File) error {
	scanner := bufio.NewScanner(file)
	var assignmentPairs []assignmentPair
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")
		assignmentPairs = append(assignmentPairs, assignmentPair{
			left:  newAssignment(pairs[0]),
			right: newAssignment(pairs[1]),
		})
	}
	d.assignmentPairs = assignmentPairs

	return nil
}
