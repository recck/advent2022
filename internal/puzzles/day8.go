package puzzles

import (
	"bufio"
	"image"
	"os"
)

var adjacentPoints = []image.Point{
	{0, -1}, // top
	{1, 0},  // right
	{-1, 0}, // left
	{0, 1},  // bottom
}

type Day8 struct {
	trees map[image.Point]int
}

func (d *Day8) Part1() any {
	visible := 0
	for point, treeHeight := range d.trees {
	nextPoint:
		for _, adjacentPoint := range adjacentPoints {
		nextAdjPoint:
			for i := 1; ; i++ {
				if existingTree, ok := d.trees[point.Add(adjacentPoint.Mul(i))]; !ok {
					visible++
					break nextPoint // mark as visible, move onto next tree
				} else if existingTree >= treeHeight {
					break nextAdjPoint // not visible in this direction, try the next direction
				}
			}
		}
	}

	return visible
}

func (d *Day8) Part2() any {
	score := 0
	bVal := map[bool]int{true: 0, false: 1}

	for point, treeHeight := range d.trees {
		curScore := 1
		for _, adjacentPoint := range adjacentPoints {
			for i := 1; ; i++ {
				if existingTree, ok := d.trees[point.Add(adjacentPoint.Mul(i))]; !ok || existingTree >= treeHeight {
					curScore *= i - bVal[ok]
					break
				}
			}
		}

		if curScore > score {
			score = curScore
		}
	}

	return score
}

func (d *Day8) LoadInput(file *os.File) error {
	scanner := bufio.NewScanner(file)
	trees := make(map[image.Point]int)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			trees[image.Point{X: x, Y: y}] = int(r - '0')
		}
		y++
	}

	d.trees = trees

	return nil
}
