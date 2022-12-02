package puzzles

import (
	"bufio"
	"os"
	"strings"
)

type NormalMove int

const (
	Rock NormalMove = iota
	Paper
	Scissor
)

type ExpectedResult int

const (
	Lose ExpectedResult = iota
	Draw
	Win
)

var letterMap = map[string]NormalMove{
	"A": Rock,
	"B": Paper,
	"C": Scissor,
	"X": Rock,
	"Y": Paper,
	"Z": Scissor,
}

var scoreMap = map[NormalMove]int{
	Rock:    1,
	Paper:   2,
	Scissor: 3,
}

var expectedResultMap = map[NormalMove]ExpectedResult{
	Rock:    Lose,
	Paper:   Draw,
	Scissor: Win,
}

type move struct {
	OpMove, MyMove NormalMove
}

func (m move) getPart1Score() int {
	baseScore := scoreMap[m.MyMove]
	// draw
	if m.OpMove == m.MyMove {
		return 3 + baseScore
	}

	// opponent wins, Rock beats Scissor, Paper covers Rock, Scissor cuts Paper
	loss := m.OpMove == Rock && m.MyMove == Scissor ||
		m.OpMove == Paper && m.MyMove == Rock ||
		m.OpMove == Scissor && m.MyMove == Paper

	if loss {
		return baseScore
	}

	return 6 + baseScore
}

func (m move) getPart2Score() int {
	expectedResult := expectedResultMap[m.MyMove]
	baseScore := 0
	myMoveScore := 0

	switch expectedResult {
	case Draw:
		baseScore = 3
		myMoveScore = scoreMap[m.OpMove]
	case Lose:
		baseScore = 0
		switch m.OpMove {
		case Rock:
			myMoveScore = scoreMap[Scissor]
		case Paper:
			myMoveScore = scoreMap[Rock]
		case Scissor:
			myMoveScore = scoreMap[Paper]
		}
	case Win:
		baseScore = 6
		switch m.OpMove {
		case Rock:
			myMoveScore = scoreMap[Paper]
		case Paper:
			myMoveScore = scoreMap[Scissor]
		case Scissor:
			myMoveScore = scoreMap[Rock]
		}
	}

	return baseScore + myMoveScore
}

type Day2 struct {
	moves []move
}

func (d *Day2) addMove(newMove move) {
	d.moves = append(d.moves, newMove)
}

func (d *Day2) Part1() any {
	sum := 0
	for i := range d.moves {
		sum += d.moves[i].getPart1Score()
	}
	return sum
}

func (d *Day2) Part2() any {
	sum := 0
	for i := range d.moves {
		sum += d.moves[i].getPart2Score()
	}
	return sum
}

func (d *Day2) LoadInput(file *os.File) error {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, " ")
		newMove := move{
			OpMove: letterMap[chunks[0]],
			MyMove: letterMap[chunks[1]],
		}
		d.addMove(newMove)
	}

	return nil
}
