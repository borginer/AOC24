package main

import (
	"bufio"
	"fmt"
	"os"
)

type boardDim struct {
	width, height int
}

var XMAS string = "XMAS"

func (bd boardDim) isOnBoard(i, j int) bool {
	return i >= 0 && j >= 0 && i < bd.height && j < bd.width
}

func checkDir(board []string, i, j, x, y int) bool {
	for step := range len(XMAS) {
		if board[i+x*step][j+y*step] != XMAS[step] {
			return false
		}
	}
	return true
}

func countCellXams(board []string, bd boardDim, i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 || !bd.isOnBoard(i+(len(XMAS)-1)*x, j+(len(XMAS)-1)*y) {
				continue
			}
			if checkDir(board, i, j, x, y) {
				count++
			}
		}
	}
	return count
}

func countXMAS(board []string) int {
	if len(board) == 0 {
		return 0
	}
	count := 0
	bd := boardDim{len(board[0]) - 1, len(board)}
	for i := range len(board) {
		for j := range len(board[0]) {
			count += countCellXams(board, bd, i, j)
		}
	}
	return count
}

func checkCellDoubleMAS(board []string, i, j int) bool {
	if board[i][j] != 'A' {
		return false
	}
	topLeft, topRight := board[i-1][j-1], board[i-1][j+1]
	botLeft, botRight := board[i+1][j-1], board[i+1][j+1]

	main := (topLeft == 'M' && botRight == 'S') || (topLeft == 'S' && botRight == 'M')
	secondary := (topRight == 'M' && botLeft == 'S') || (topRight == 'S' && botLeft == 'M')
	return main && secondary
}

func countDoubleMAS(board []string) int {
	if len(board) == 0 {
		return 0
	}
	count := 0
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-2; j++ {
			if checkCellDoubleMAS(board, i, j) {
				count++
			}
		}
	}
	return count
}

func main() {
	board := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		board = append(board, line)
	}

	fmt.Println("part 1:", countXMAS(board))
	fmt.Println("part 2:", countDoubleMAS(board))
}
