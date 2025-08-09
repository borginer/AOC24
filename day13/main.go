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
	CostA      = 3
	CostB      = 1
	MaxPresses = 100
	Skew       = 10000000000000
)

type Button struct {
	x, y, cost int64
}
type Pos struct {
	x, y int64
}

func parseXYPair(text, sep string) (x, y int64) {
	input := strings.Split(text, ":")[1]
	stats := strings.Split(input, ",")
	xStat := strings.Split(stats[0], sep)[1]
	yStat := strings.Split(stats[1], sep)[1]
	var err error
	x, err = strconv.ParseInt(xStat, 10, 64)
	if err != nil {
		log.Fatal("parsing button f")
	}
	y, err = strconv.ParseInt(yStat, 10, 64)
	if err != nil {
		log.Fatal("parsing button f")
	}
	return x, y
}

func parseMachineStats(s *bufio.Scanner) (A, B Button, Target Pos, ok bool) {
	if s.Scan() {
		buttonAtext := s.Text()
		x, y := parseXYPair(buttonAtext, "+")
		A = Button{x, y, CostA}
		if s.Scan() {
			buttonBtext := s.Text()
			x, y = parseXYPair(buttonBtext, "+")
			B = Button{x, y, CostB}

			if s.Scan() {
				prizeText := s.Text()
				x, y = parseXYPair(prizeText, "=")
				Target = Pos{x, y}
				s.Scan()
				return A, B, Target, true
			}
		}
	}
	return Button{}, Button{}, Pos{}, false
}

func calcMinTokenCost(A, B Button, tar Pos) int64 {
	x1, x2, y1, y2 := float64(A.x), float64(B.x), float64(A.y), float64(B.y)
	det := x1*y2 - x2*y1
	if det == 0 {
		log.Fatal("only inversible matrixes supported")
	}
	tx, ty := float64(tar.x), float64(tar.y)
	b1 := (y2*tx - x2*ty) / det
	b2 := (x1*ty - y1*tx) / det
	if b1 == float64(int64(b1)) && b2 == float64(int64(b2)) {
		return int64(b1)*A.cost + int64(b2)*B.cost
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalCost := int64(0)
	totalCostSkew := int64(0)
	for {
		A, B, tar, ok := parseMachineStats(scanner)
		if !ok {
			break
		}
		totalCost += calcMinTokenCost(A, B, tar)
		totalCostSkew += calcMinTokenCost(A, B, Pos{tar.x + Skew, tar.y + Skew})
	}
	fmt.Println("part 1", totalCost)
	fmt.Println("part 2", totalCostSkew)
}
