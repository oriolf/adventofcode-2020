package main

import (
	"fmt"
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	space := getInitialSpace(lines)
	for i := 0; i < 6; i++ {
		space = performCycle(space)
	}

	return countActive(space)
}

func solve2(lines []string) interface{} {
	space := getInitialSpace(lines)
	n := len(space)
	hyperSpace := [][][][]string{newSpace(n), space, newSpace(n)}
	for i := 0; i < 6; i++ {
		hyperSpace = performHyperCycle(hyperSpace)
	}

	return countHyperActive(hyperSpace)
}

func getInitialSpace(lines []string) [][][]string {
	plane := [][]string{}
	for _, l := range lines {
		line := strings.Split(l, "")
		if len(line)%2 == 0 {
			line = append(line, ".")
		}
		plane = append(plane, line)
	}

	if len(plane)%2 == 0 {
		plane = append(plane, newLine(len(plane[0])))
	}

	n := len(plane)
	extra := len(plane) - 1
	space := [][][]string{plane}
	for i := 0; i < extra/2; i++ {
		space = append([][][]string{newPlane(n)}, space...)
		space = append(space, newPlane(n))
	}
	return space
}

func printSpace(space [][][]string) {
	for i, plane := range space {
		fmt.Println("z = ", i)
		printPlane(plane)
	}
}

func printPlane(plane [][]string) {
	for _, line := range plane {
		fmt.Println(strings.Join(line, ""))
	}
}

func performCycle(space [][][]string) [][][]string {
	space = expandSpace(space)
	counts := countActives(space)
	for i, plane := range counts {
		for j, line := range plane {
			for k, count := range line {
				space[i+1][j+1][k+1] = updateActive(space[i+1][j+1][k+1], count)
			}
		}
	}

	return space
}

func performHyperCycle(hyperSpace [][][][]string) [][][][]string {
	hyperSpace = expandHyperSpace(hyperSpace)
	counts := countHyperActives(hyperSpace)
	for n, space := range counts {
		for i, plane := range space {
			for j, line := range plane {
				for k, count := range line {
					hyperSpace[n+1][i+1][j+1][k+1] = updateActive(hyperSpace[n+1][i+1][j+1][k+1], count)
				}
			}
		}
	}

	return hyperSpace
}

func countActives(space [][][]string) (counts [][][]int) {
	for _, plane := range space[1 : len(space)-1] {
		var p [][]int
		for _, line := range plane[1 : len(plane)-1] {
			var l []int
			for range line[1 : len(line)-1] {
				l = append(l, 0)
			}
			p = append(p, l)
		}
		counts = append(counts, p)
	}

	for i, plane := range space[1 : len(space)-1] {
		for j, line := range plane[1 : len(plane)-1] {
			for k := range line[1 : len(line)-1] {
				counts[i][j][k] = countActivesSingle(space, i+1, j+1, k+1)
			}
		}
	}

	return counts
}

func countActivesSingle(space [][][]string, i, j, k int) (count int) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x != 0 || y != 0 || z != 0 {
					if space[i+x][j+y][k+z] == "#" {
						count++
					}
				}
			}
		}
	}

	return count
}

func countHyperActives(hyperSpace [][][][]string) (counts [][][][]int) {
	for _, space := range hyperSpace[1 : len(hyperSpace)-1] {
		var s [][][]int
		for _, plane := range space[1 : len(space)-1] {
			var p [][]int
			for _, line := range plane[1 : len(plane)-1] {
				var l []int
				for range line[1 : len(line)-1] {
					l = append(l, 0)
				}
				p = append(p, l)
			}
			s = append(s, p)
		}
		counts = append(counts, s)
	}

	for n, space := range hyperSpace[1 : len(hyperSpace)-1] {
		for i, plane := range space[1 : len(space)-1] {
			for j, line := range plane[1 : len(plane)-1] {
				for k := range line[1 : len(line)-1] {
					counts[n][i][j][k] = countHyperActivesSingle(hyperSpace, n+1, i+1, j+1, k+1)
				}
			}
		}
	}

	return counts
}

func countHyperActivesSingle(hyperSpace [][][][]string, n, i, j, k int) (count int) {
	for t := -1; t <= 1; t++ {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					if t != 0 || x != 0 || y != 0 || z != 0 {
						if hyperSpace[n+t][i+x][j+y][k+z] == "#" {
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func updateActive(current string, count int) string {
	if current == "#" {
		if count == 2 || count == 3 {
			return "#"
		}
		return "."
	}

	if count == 3 {
		return "#"
	}

	return "."
}

func hyperSpaceDimensions(space [][][][]string) (int, int, int, int) {
	return len(space), len(space[0]), len(space[0][0]), len(space[0][0][0])
}

func printHyperDimensions(hyper [][][][]string) {
	fmt.Println(len(hyper))
	for _, space := range hyper {
		fmt.Println("  ", len(space))
		for _, plane := range space {
			fmt.Println("    ", len(plane))
			for _, line := range plane {
				fmt.Println("      ", len(line))
			}
		}
	}
}

func spaceDimensions(space [][][]string) (int, int, int) {
	return len(space), len(space[0]), len(space[0][0])
}

func expandHyperSpace(hyperSpace [][][][]string) [][][][]string {
	newHyper := [][][][]string{newSpace(len(hyperSpace[0]) + 2)}
	for _, space := range hyperSpace {
		newHyper = append(newHyper, expandSpace(space))
	}
	return append(newHyper, newSpace(len(hyperSpace[0])+2))
}

func expandSpace(space [][][]string) [][][]string {
	newSpace := [][][]string{newPlane(len(space[0]) + 2)}
	for _, plane := range space {
		newSpace = append(newSpace, expandPlane(plane))
	}

	return append(newSpace, newPlane(len(space[0])+2))
}

func expandPlane(plane [][]string) [][]string {
	newPlane := [][]string{newLine(len(plane[0]) + 2)}
	for _, line := range plane {
		newPlane = append(newPlane, expandLine(line))
	}

	return append(newPlane, newLine(len(plane[0])+2))
}

func expandLine(line []string) []string {
	newLine := append([]string{"."}, line...)
	return append(newLine, ".")
}

func newSpace(length int) (space [][][]string) {
	for i := 0; i < length; i++ {
		space = append(space, newPlane(length))
	}
	return space
}

func newPlane(length int) (plane [][]string) {
	for i := 0; i < length; i++ {
		plane = append(plane, newLine(length))
	}
	return plane
}

func newLine(length int) (line []string) {
	for i := 0; i < length; i++ {
		line = append(line, ".")
	}
	return line
}

func countHyperActive(hyperSpace [][][][]string) (count int) {
	for _, space := range hyperSpace {
		count += countActive(space)
	}
	return count
}

func countActive(space [][][]string) (count int) {
	for _, x := range space {
		for _, y := range x {
			for _, z := range y {
				if z == "#" {
					count++
				}
			}
		}
	}

	return count
}
