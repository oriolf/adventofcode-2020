package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	player1, player2 := parsePlayers(lines)
	player1, player2 = playGame(player1, player2)
	return computeScore(winner(player1, player2))
}

func solve2(lines []string) interface{} {
	player1, player2 := parsePlayers(lines)
	player1, player2, _ = playGame2(player1, player2, nil)
	return computeScore(winner(player1, player2))
}

func winner(p1, p2 []int) []int {
	if len(p2) > 0 {
		return p2
	}
	return p1
}

func parsePlayers(lines []string) (p1, p2 []int) {
	var i int
	for i = 1; lines[i] != ""; i++ {
		p1 = append(p1, util.ParseInt(lines[i]))
	}

	i += 2
	for ; i < len(lines); i++ {
		p2 = append(p2, util.ParseInt(lines[i]))
	}

	return p1, p2
}

func playGame(p1, p2 []int) ([]int, []int) {
	for len(p1) > 0 && len(p2) > 0 {
		x1, x2 := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]
		if x1 > x2 {
			p1 = append(p1, x1, x2)
		} else if x2 > x1 {
			p2 = append(p2, x2, x1)
		} else {
			panic("impossible!")
		}
	}

	return p1, p2
}

func computeScore(in []int) (score int) {
	mult := len(in)
	for _, x := range in {
		score += mult * x
		mult--
	}

	return score
}

func playGame2(p1, p2 []int, previous map[int][][]int) ([]int, []int, int) {
	if previous == nil {
		previous = make(map[int][][]int)
	}

	for len(p1) > 0 && len(p2) > 0 {
		if alreadySeen(p1, p2, previous) {
			return p1, p2, 0
		}
		previous[0] = append(previous[0], p1)
		previous[1] = append(previous[1], p2)

		x1, x2 := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]
		var winner int
		if len(p1) >= x1 && len(p2) >= x2 {
			_, _, winner = playGame2(Copy(p1[:x1]), Copy(p2[:x2]), nil)
		} else {
			if x1 > x2 {
				winner = 0
			} else if x2 > x1 {
				winner = 1
			} else {
				panic("impossible!")
			}
		}

		if winner == 0 {
			p1 = append(p1, x1, x2)
		} else {
			p2 = append(p2, x2, x1)
		}
	}

	var winner int
	if len(p1) == 0 {
		winner = 1
	}

	return p1, p2, winner
}

func Copy(in []int) (out []int) {
	return append([]int{}, in...)
}

func alreadySeen(p1, p2 []int, previous map[int][][]int) bool {
	for i := range previous[0] {
		if Equal(previous[0][i], p1) && Equal(previous[1][i], p2) {
			return true
		}
	}
	return false
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}

	return true
}
