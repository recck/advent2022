package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var monkeyLcm int

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type monkey struct {
	items                                 []int
	operation                             func(int) int
	test                                  int
	truthMonkey, falseMonkey, inspections int
}

func (m *monkey) addItem(item int) {
	m.items = append(m.items, item)
}

func newMonkey() monkey {
	return monkey{
		items:       []int{},
		operation:   nil,
		test:        0,
		truthMonkey: 0,
		falseMonkey: 0,
		inspections: 0,
	}
}

type Day11 struct {
	monkeys []monkey
}

func (d *Day11) performInspection(monkeyId int, damaged bool) {
	curMonkey := d.monkeys[monkeyId]
	monkeyItems := curMonkey.items
	d.monkeys[monkeyId].items = []int{}
	for i := range monkeyItems {
		worryLevel := curMonkey.operation(monkeyItems[i])
		if damaged {
			worryLevel /= 3
		} else {
			worryLevel %= monkeyLcm
		}
		if worryLevel%curMonkey.test == 0 {
			d.monkeys[curMonkey.truthMonkey].addItem(worryLevel)
		} else {
			d.monkeys[curMonkey.falseMonkey].addItem(worryLevel)
		}
		d.monkeys[monkeyId].inspections++
	}
}

func (d *Day11) performInspections(numRounds int, damaged bool) {
	for round := 0; round < numRounds; round++ {
		for monkeyId := range d.monkeys {
			d.performInspection(monkeyId, damaged)
		}
	}
}

func (d *Day11) getMonkeyBusiness() int {
	var inspectionCounts []int
	for i := range d.monkeys {
		inspectionCounts = append(inspectionCounts, d.monkeys[i].inspections)
	}

	sort.Ints(inspectionCounts)

	return inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2]
}

func (d *Day11) Part1() any {
	monkeyCopy := make([]monkey, len(d.monkeys))
	copy(monkeyCopy, d.monkeys)
	d.performInspections(20, true)

	monkeyBusiness := d.getMonkeyBusiness()
	d.monkeys = monkeyCopy
	return monkeyBusiness
}

func (d *Day11) Part2() any {
	d.performInspections(10000, false)
	return d.getMonkeyBusiness()
}

func (d *Day11) LoadInput(file *os.File) error {
	var monkeys []monkey
	curMonkeyId := 0
	curLine := 0
	curMonkey := newMonkey()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			curMonkeyId++
			curLine = 0
			monkeys = append(monkeys, curMonkey)
			curMonkey = newMonkey()
			continue
		}

		chunks := strings.Split(line, ":")

		switch curLine {
		case 1:
			items := strings.Split(strings.TrimSpace(chunks[1]), ",")
			var intItems []int
			for i := range items {
				intItem, _ := strconv.Atoi(strings.TrimSpace(items[i]))
				intItems = append(intItems, intItem)
			}
			curMonkey.items = intItems
		case 2:
			fields := strings.Fields(chunks[1])
			operation := fields[3]
			modifier := fields[4]
			if modifier == "old" {
				curMonkey.operation = func(v int) int {
					return v * v
				}
			} else {
				modifierInt, _ := strconv.Atoi(modifier)
				if operation == "+" {
					curMonkey.operation = func(v int) int {
						return v + modifierInt
					}
				} else {
					curMonkey.operation = func(v int) int {
						return v * modifierInt
					}
				}
			}
		case 3:
			var modifier int
			fmt.Sscanf(line, "Test: divisible by %d", &modifier)
			curMonkey.test = modifier
		case 4:
			var trueMonkey int
			fmt.Sscanf(line, "If true: throw to monkey %d", &trueMonkey)
			curMonkey.truthMonkey = trueMonkey
		case 5:
			var falseMonkey int
			fmt.Sscanf(line, "If false: throw to monkey %d", &falseMonkey)
			curMonkey.falseMonkey = falseMonkey
		}
		curLine++
	}
	monkeys = append(monkeys, curMonkey)

	var tests []int
	for i := range monkeys {
		tests = append(tests, monkeys[i].test)
	}
	monkeyLcm = LCM(tests[0], tests[1], tests[2:]...)

	d.monkeys = monkeys
	return nil
}
