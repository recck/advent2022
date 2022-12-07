package puzzles

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct {
	fileTree map[string]int
}

func (d *Day7) getDirSum(dir string) int {
	sum := 0
	for dirs := range d.fileTree {
		if strings.Contains(dirs, dir) {
			sum += d.fileTree[dirs]
		}
	}

	return sum
}

func (d *Day7) Part1() any {
	sum := 0
	for dir := range d.fileTree {
		fileSum := d.getDirSum(dir)

		if fileSum <= 100000 {
			sum += fileSum
		}
	}

	return sum
}

func (d *Day7) Part2() any {
	totalSize := 70000000
	neededSpace := 30000000
	usedSpace := d.getDirSum("/")
	unusedSpace := totalSize - usedSpace

	var dirSizes []int
	for dir := range d.fileTree {
		dirSizes = append(dirSizes, d.getDirSum(dir))
	}
	sort.Ints(dirSizes)

	for i := 0; i < len(dirSizes); i++ {
		if unusedSpace+dirSizes[i] >= neededSpace {
			return dirSizes[i]
		}
	}

	return 0
}

func (d *Day7) LoadInput(f *os.File) error {
	var curDir string
	var dirs []string

	tree := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, " ")
		if strings.IndexRune(line, '$') == 0 {
			if line == "$ ls" {
				continue
			}

			nextDir := chunks[2]

			if nextDir == ".." {
				curDir = dirs[len(dirs)-2]
				dirs = dirs[:len(dirs)-1]
			} else {
				curDir = nextDir
				dirs = append(dirs, curDir)
			}
		} else {
			// we must be ls-ing
			if chunks[0] == "dir" {
				dirString := strings.Join(append(dirs, chunks[1]), "_")
				if _, ok := tree[dirString]; !ok {
					tree[dirString] = 0
				}
			} else {
				fileSize, _ := strconv.Atoi(chunks[0])
				dirString := strings.Join(dirs, "_")
				if existingDir, ok := tree[dirString]; !ok {
					tree[dirString] = fileSize
				} else {
					tree[dirString] = existingDir + fileSize
				}
			}
		}
	}

	d.fileTree = tree

	return nil
}
