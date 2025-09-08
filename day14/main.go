package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	BoardWidth  = 101
	BoardHeight = 103
	Time        = 100
)

type Vec struct {
	x, y int64
}

func parseNumberPair(pair []string) Vec {
	var ret Vec
	var err error
	ret.x, err = strconv.ParseInt(pair[0], 10, 64)
	if err != nil {
		log.Fatal("unlucky")
	}
	ret.y, err = strconv.ParseInt(pair[1], 10, 64)
	if err != nil {
		log.Fatal("unlucky2")
	}
	return ret
}

func parseLine(line string) (Vec, Vec) {
	parts := strings.Fields(line)
	left, right := parts[0], parts[1]
	leftVecStr := strings.Split(left, "=")[1]
	rightVecStr := strings.Split(right, "=")[1]
	leftIndexes := strings.Split(leftVecStr, ",")
	rightIndexes := strings.Split(rightVecStr, ",")
	var p, v Vec
	p = parseNumberPair(leftIndexes)
	v = parseNumberPair(rightIndexes)
	return p, v
}

func getEndingPosition(p, v Vec, time int64) Vec {
	endX := (p.x + time*v.x) % BoardWidth
	if endX < 0 {
		endX += BoardWidth
	}
	endY := (p.y + time*v.y) % BoardHeight
	if endY < 0 {
		endY += BoardHeight
	}
	return Vec{endX, endY}
}

func calculateSafetyFactor(robots map[Vec]int) int {
	var tlCount, trCount, blCount, brCount int
	for i := range int64(BoardWidth) {
		for j := range int64(BoardHeight) {
			if i < BoardWidth/2 && j < BoardHeight/2 {
				tlCount += robots[Vec{i, j}]
			} else if i > BoardWidth/2 && j < BoardHeight/2 {
				trCount += robots[Vec{i, j}]
			} else if i < BoardWidth/2 && j > BoardHeight/2 {
				blCount += robots[Vec{i, j}]
			} else if i > BoardWidth/2 && j > BoardHeight/2 {
				brCount += robots[Vec{i, j}]
			}
		}
	}
	return tlCount * trCount * blCount * brCount
}

func calcClatter(robots map[Vec]int) int {
	clatter := 0
	for i := range int64(BoardWidth) {
		for j := range int64(BoardHeight) {
			if robots[Vec{i, j}] >= 1 {
				clatter += robots[Vec{i, j + 1}]
				clatter += robots[Vec{i + 1, j}]
				clatter += robots[Vec{i - 1, j}]
				clatter += robots[Vec{i, j - 1}]
			}
		}
	}
	return clatter
}

func printBoard(robots map[Vec]int, time int64) {
	if calcClatter(robots) > 500 {
		for i := range int64(BoardWidth) {
			for j := range int64(BoardHeight) {
				if robots[Vec{i, j}] == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(robots[Vec{i, j}])
				}
			}
			fmt.Println()
		}
		fmt.Println(time, "\n\n\n")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var p, v Vec
	robotsP1 := make(map[Vec]int)
	robotsP2 := make(map[Vec]int)
	var pSlice, vSlice []Vec
	for scanner.Scan() {
		p, v = parseLine(scanner.Text())
		robotsP1[getEndingPosition(p, v, Time)]++
		pSlice = append(pSlice, p)
		vSlice = append(vSlice, v)
	}
	for time := range int64(10000) {
		for i, p := range pSlice {
			robotsP2[getEndingPosition(p, vSlice[i], time)]++
		}
		printBoard(robotsP2, time)
		clear(robotsP2)
	}
	fmt.Println("Part 1: ", calculateSafetyFactor(robotsP1))
}
