package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	position, lastDirection := [2]int{0, 0}, "E"
	for _, l := range lines {
		direction, value := l[0:1], util.ParseInt(l[1:])
		position, lastDirection = updatePosition(position, lastDirection, direction, value)
	}

	return util.Abs(position[0]) + util.Abs(position[1])
}

func solve2(lines []string) interface{} {
	waypoint, position := [2]int{10, 1}, [2]int{0, 0}
	for _, l := range lines {
		direction, value := l[0:1], util.ParseInt(l[1:])
		waypoint, position = updatePosition2(waypoint, position, direction, value)
	}

	return util.Abs(position[0]) + util.Abs(position[1])
}

func updatePosition2(waypoint, position [2]int, direction string, value int) ([2]int, [2]int) {
	if direction == "F" {
		position[0] += waypoint[0] * value
		position[1] += waypoint[1] * value
		return waypoint, position
	}

	switch direction {
	case "N", "E", "S", "W":
		waypoint = updatePositionStraight(waypoint, direction, value)
	case "R":
		waypoint = rotateWaypointRight(waypoint, value/90)
	case "L":
		waypoint = rotateWaypointLeft(waypoint, value/90)
	}

	return waypoint, position
}

func rotateWaypointRight(w [2]int, steps int) [2]int {
	for i := 0; i < steps; i++ {
		w[0], w[1] = w[1], -w[0]
	}
	return w
}

func rotateWaypointLeft(w [2]int, steps int) [2]int {
	for i := 0; i < steps; i++ {
		w[0], w[1] = -w[1], w[0]
	}
	return w
}

func updatePosition(pos [2]int, lastDirection, direction string, value int) ([2]int, string) {
	if direction == "R" {
		return pos, rotateRight(lastDirection, value/90)
	} else if direction == "L" {
		return pos, rotateLeft(lastDirection, value/90)
	}

	if direction == "F" {
		direction = lastDirection
	}

	pos = updatePositionStraight(pos, direction, value)

	return pos, lastDirection
}

func updatePositionStraight(pos [2]int, direction string, value int) [2]int {
	switch direction {
	case "N":
		pos[1] += value
	case "E":
		pos[0] += value
	case "S":
		pos[1] -= value
	case "W":
		pos[0] -= value
	default:
		panic("unknown direction")
	}
	return pos
}

var (
	directionsRight = []string{"N", "E", "S", "W"}
	directionsLeft  = []string{"N", "W", "S", "E"}
)

func rotateRight(direction string, steps int) string {
	i := util.IndexOf(direction, directionsRight)
	return directionsRight[(i+steps)%4]
}

func rotateLeft(direction string, steps int) string {
	i := util.IndexOf(direction, directionsLeft)
	return directionsLeft[(i+steps)%4]
}
