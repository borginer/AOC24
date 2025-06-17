package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left, right []int
	leftMap, rightMap := make(map[int]int), make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "   ")
		lnum, _ := strconv.Atoi(res[0])
		leftMap[lnum]++
		left = append(left, lnum)
		rnum, _ := strconv.Atoi(res[1])
		rightMap[rnum]++
		right = append(right, rnum)
	}

	sort.Ints(left)
	sort.Ints(right)

	diff := int64(0)
	for i := range left {
		dist := max(right[i], left[i]) - min(right[i], left[i])
		diff += int64(dist)
	}

	fmt.Println("list diff:", diff) // part 1

	similarity := int64(0)
	for num, val := range leftMap {
		similarity += int64(num * val * rightMap[num])
	}

	fmt.Println("list similarity:", similarity)
}
