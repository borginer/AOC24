package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

const (
	STEP_MARK = 'X'
	OBSTACLE  = '#'
)

type Square struct {
	i, j int
}

var UP Square = Square{-1, 0}
var RIGHT Square = Square{0, 1}
var DOWN Square = Square{1, 0}
var LEFT Square = Square{0, -1}

var rightTurn = map[Square]Square{
	{-1, 0}: {0, 1},
	{0, 1}:  {1, 0},
	{1, 0}:  {0, -1},
	{0, -1}: {-1, 0},
}

func (s Square) add(other Square) Square {
	s.i += other.i
	s.j += other.j
	return s
}

type Lab [][]byte

func (l Lab) contains(s Square) bool {
	return s.i >= 0 && s.j >= 0 && s.i < len(l) && s.j < len(l[0])
}

func (l Lab) get(s Square) byte {
	return l[s.i][s.j]
}

func (l Lab) set(s Square, val byte) {
	l[s.i][s.j] = val
}

func findGuard(lab Lab) (Square, bool) {
	for i, line := range lab {
		for j, s := range line {
			if s == '^' {
				return Square{i, j}, true
			}
		}
	}
	return Square{}, false
}

func countSteps(lab Lab) (steps int) {
	for _, line := range lab {
		for _, s := range line {
			if s == STEP_MARK {
				steps++
			}
		}
	}
	return
}

func walkInLab(lab Lab, g Square) [][]byte {
	dir := UP
	for lab.contains(g) {
		lab.set(g, STEP_MARK)
		nextStep := g.add(dir)
		if lab.contains(nextStep) && lab.get(nextStep) == OBSTACLE {
			dir = rightTurn[dir]
		}
		g = g.add(dir)
	}
	return lab
}

func isInfLoop(lab Lab, g, dir, start Square) bool {
	pastDirs := make(map[Square][]Square)
	obsPos := g.add(dir)
	g = start
	dir = UP
	if lab.get(obsPos) == OBSTACLE {
		return false
	}
	lab.set(obsPos, OBSTACLE)

	for lab.contains(g) {
		if slices.Contains(pastDirs[g], dir) {
			lab.set(obsPos, STEP_MARK)
			return true
		}
		pastDirs[g] = append(pastDirs[g], dir)
		nextStep := g.add(dir)
		if lab.contains(nextStep) && lab.get(nextStep) == OBSTACLE {
			dir = rightTurn[dir]
		} else {
			g = g.add(dir)
		}
	}

	lab.set(obsPos, STEP_MARK)
	return false
}

func countInfiniteLoops(lab Lab, g Square) int {
	obsts := make(map[Square]bool)
	dir := UP
	start := g
	for lab.contains(g) {
		nextStep := g.add(dir)
		if lab.contains(nextStep) {
			if lab.get(nextStep) == OBSTACLE {
				dir = rightTurn[dir]
				continue
			}
			if isInfLoop(lab, g, dir, start) {
				obsts[nextStep] = true
			}
			g = g.add(dir)
		} else {
			break
		}

	}
	return len(obsts)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lab Lab
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lab = append(lab, line)
	}

	start, ok := findGuard(lab)
	if !ok {
		log.Fatal("no guard :O")
	}
	lab = walkInLab(lab, start)
	fmt.Println("part 1:", countSteps(lab))
	fmt.Println("part 2:", countInfiniteLoops(lab, start))
}
