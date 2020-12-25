package main

import (
	"fmt"
	"github.com/oriolf/adventofcode2020/util"
	"math"
	"strings"
)

var monster [][]string

func init() {
	var m = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

	for _, l := range strings.Split(m, "\n") {
		monster = append(monster, strings.Split(l, ""))
	}
}

type tile struct {
	contents [][]string
	id       int
}

type match struct {
	first   int
	second  int
	flipped bool
}

func main() {
	util.Solve(solve1, solve2)
}

func print(args ...interface{}) {
	fmt.Println(args...)
}

func solve1(lines []string) interface{} {
	tiles := parseTiles(lines)

	corners := getCorners(tiles)

	mult := 1
	for _, c := range corners {
		mult *= c.id
	}

	return mult
}

func getCorners(tiles []tile) (corners []tile) {
	for _, t := range tiles {
		count := countMatchingEdges(t, tiles)
		switch count {
		case 2:
			corners = append(corners, t)
		case 3, 4:
		default:
			panic(fmt.Sprintf("unexpected number of matches: %v", count))
		}
	}

	return corners
}

func solve2(lines []string) interface{} {
	image := reconstructImage(lines)
	finalImage := getFinalImage(image)
	monsterMap := getMonsterMap(finalImage)

	var count int
	for i, row := range finalImage {
		for j, x := range row {
			if x == "#" && !monsterMap[i][j] {
				count++
			}
		}
	}

	return count
}

func getMonsterMap(image [][]string) [][]bool {
	var m [][]bool
	for range image {
		var r []bool
		for range image {
			r = append(r, false)
		}
		m = append(m, r)
	}

	return markMonsters(image, m)
}

func markMonsters(image [][]string, m [][]bool) [][]bool {
	for i := 0; i < len(image)-len(monster); i++ {
		for j := 0; j < len(image[0])-len(monster[0]); j++ {
			if hasMonsterAt(image, i, j) {
				m = markMonsterAt(m, i, j)
			}
		}
	}

	return m
}

func markMonsterAt(m [][]bool, i, j int) [][]bool {
	for x, row := range monster {
		for y := range row {
			if monster[x][y] == "#" {
				m[i+x][j+y] = true
			}
		}
	}

	return m
}

func hasMonsters(image [][]string) bool {
	for i := 0; i < len(image)-len(monster); i++ {
		for j := 0; j < len(image[0])-len(monster[0]); j++ {
			if hasMonsterAt(image, i, j) {
				return true
			}
		}
	}

	return false
}

func hasMonsterAt(image [][]string, i, j int) bool {
	for x, row := range monster {
		for y := range row {
			if monster[x][y] == "#" && image[i+x][j+y] != "#" {
				return false
			}
		}
	}

	return true
}

func getFinalImage(image [][]tile) (out [][]string) {
	for _, row := range image {
		out = append(out, getRowPixels(row)...)
	}

	for !hasMonsters(out) {
		if hasMonsters(flipMatrixRows(out)) {
			out = flipMatrixRows(out)
		} else if hasMonsters(flipMatrixColumns(out)) {
			out = flipMatrixColumns(out)
		} else {
			out = rotateMatrix(out)
		}
	}

	return out
}

func getRowPixels(row []tile) (out [][]string) {
	for i := 1; i < len(row[0].contents)-1; i++ {
		var r []string
		for _, t := range row {
			n := len(t.contents[i])
			r = append(r, t.contents[i][1:n-1]...)
		}
		out = append(out, r)
	}
	return out
}

func reconstructImage(lines []string) (image [][]tile) {
	tiles := parseTiles(lines)
	width := int(math.Sqrt(float64(len(tiles))))

	lastAdded := getFirstCorner(tiles)
	row := []tile{lastAdded}
	tiles = removeTile(tiles, lastAdded)
	goingRight := true
LOOP:
	for len(tiles) > 0 {
		for _, t := range tiles {
			if goingRight && matchEdge(lastAdded.right(), t.edges()) {
				t = convertMatchLeft(lastAdded.right(), t)
				lastAdded = t
				row = append(row, t)
				if len(row) == width {
					goingRight = false
					image = append(image, row)
					lastAdded = row[0]
					row = nil
				}

				tiles = removeTile(tiles, t)
				continue LOOP
			} else if !goingRight && matchEdge(lastAdded.bottom(), t.edges()) {
				t = convertMatchTop(lastAdded.bottom(), t)
				lastAdded = t
				row = []tile{t}
				goingRight = true
				tiles = removeTile(tiles, t)
				continue LOOP
			} else {
				continue
			}
		}

		panic("could not add any tile!")
	}

	return image
}

func getFirstCorner(tiles []tile) tile {
	first := getCorners(tiles)[0]
	for !firstCornerWellOriented(first, tiles) {
		if firstCornerWellOriented(flipRows(first), tiles) {
			return flipRows(first)
		} else if firstCornerWellOriented(flipColumns(first), tiles) {
			return flipColumns(first)
		} else {
			first = rotateTile(first)
		}
	}

	return first
}

func firstCornerWellOriented(corner tile, tiles []tile) bool {
	a := edgeExactlyMatches(corner.right(), removeTile(tiles, corner))
	b := edgeExactlyMatches(corner.bottom(), removeTile(tiles, corner))
	return a && b
}

func edgeExactlyMatches(e []string, tiles []tile) bool {
	for _, t := range tiles {
		for _, x := range t.edges() {
			if matches(e, x) || matches(e, flip(x)) {
				return true
			}
		}
	}
	return false
}

func convertMatchLeft(e []string, t tile) tile {
	for !matches(e, t.left()) {
		if matches(e, flip(t.left())) {
			return flipRows(t)
		}
		t = rotateTile(t)
	}
	return t
}

func convertMatchTop(e []string, t tile) tile {
	for !matches(e, t.top()) {
		if matches(e, flip(t.top())) {
			return flipColumns(t)
		}
		t = rotateTile(t)
	}
	return t
}

func flipRows(t tile) tile {
	return tile{id: t.id, contents: flipMatrixRows(t.contents)}
}

func flipMatrixRows(in [][]string) [][]string {
	var contents [][]string
	for i := len(in) - 1; i >= 0; i-- {
		contents = append(contents, in[i])
	}
	return contents
}

func flipColumns(t tile) tile {
	return tile{id: t.id, contents: flipMatrixColumns(t.contents)}
}

func flipMatrixColumns(in [][]string) [][]string {
	var contents [][]string
	for _, row := range in {
		contents = append(contents, flip(row))
	}
	return contents
}

func rotateTile(t tile) tile {
	return tile{id: t.id, contents: rotateMatrix(t.contents)}
}

func rotateMatrix(in [][]string) (out [][]string) {
	for j := range in {
		var r []string
		for i := len(in) - 1; i >= 0; i-- {
			r = append(r, in[i][j])
		}
		out = append(out, r)
	}

	return out
}

func imageIDs(image [][]tile) (out [][]int) {
	for _, row := range image {
		var r []int
		for _, t := range row {
			r = append(r, t.id)
		}
		out = append(out, r)
	}
	return out
}

func parseTiles(lines []string) (tiles []tile) {
	var t tile
	for _, l := range lines {
		if strings.HasPrefix(l, "Tile") {
			t.id = util.ParseInt(strings.Trim(l, "Tile :"))
		} else if l == "" {
			tiles = append(tiles, t)
			t = tile{}
		} else {
			t.contents = append(t.contents, strings.Split(l, ""))
		}
	}

	return append(tiles, t)
}

func countMatchingEdges(original tile, tiles []tile) (count int) {
	for _, t := range tiles {
		if original.id != t.id {
			count += len(matchEdges(original.edges(), t.edges()))
		}
	}

	return count
}

func matchEdges(e1, e2 [][]string) (out []match) {
	for i, x := range e1 {
		for j, y := range e2 {
			if matches(x, y) {
				out = append(out, match{first: i, second: j})
			} else if matches(x, flip(y)) {
				out = append(out, match{first: i, second: j, flipped: true})
			}
		}
	}
	return out
}

func matchEdge(e []string, ee [][]string) bool {
	for _, x := range ee {
		if matches(e, x) || matches(e, flip(x)) {
			return true
		}
	}
	return false
}

func (t tile) edges() [][]string {
	return [][]string{t.top(), t.right(), t.bottom(), t.left()}
}

func (t tile) top() []string {
	return t.contents[0]
}

func (t tile) bottom() []string {
	return t.contents[len(t.contents)-1]
}

func (t tile) left() (out []string) {
	for _, l := range t.contents {
		out = append(out, l[0])
	}
	return out
}

func (t tile) right() (out []string) {
	for _, l := range t.contents {
		out = append(out, l[len(l)-1])
	}
	return out
}

func flip(l []string) (out []string) {
	for i := len(l) - 1; i >= 0; i-- {
		out = append(out, l[i])
	}
	return out
}

func matches(a, b []string) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func removeTile(in []tile, t tile) (out []tile) {
	for _, x := range in {
		if x.id != t.id {
			out = append(out, x)
		}
	}
	return out
}
