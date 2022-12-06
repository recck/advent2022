package puzzles

import (
	"bufio"
	"os"
)

func isUnique(window []byte) bool {
	seen := make(map[byte]interface{})
	for i := range window {
		if _, ok := seen[window[i]]; ok {
			return false
		} else {
			seen[window[i]] = true
		}
	}

	return true
}

type Day6 struct {
	signal []byte
}

func (d *Day6) getFirstUniqueWindow(windowSize int) int {
	var i int
	for i = windowSize; i <= len(d.signal); i++ {
		window := d.signal[i-windowSize : i]
		if isUnique(window) {
			break
		}
	}

	return i
}

func (d *Day6) Part1() any {
	return d.getFirstUniqueWindow(4)
}

func (d *Day6) Part2() any {
	return d.getFirstUniqueWindow(14)
}

func (d *Day6) LoadInput(file *os.File) error {
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()

	if err != nil {
		return err
	}

	d.signal = line
	return nil
}
