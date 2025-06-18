package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkInRange(left int, right int) bool {
	diff := max(left, right) - min(left, right)
	if diff < 1 || diff > 3 {
		return false
	}
	return true
}

func checkMonotone(left int, right int, isIncreasing bool) bool {
	if right > left != isIncreasing {
		return false
	}
	return true
}

func checkReport(levels []int) bool {
	if len(levels) <= 1 {
		return true
	}

	isIncreasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		if !checkInRange(levels[i-1], levels[i]) ||
			!checkMonotone(levels[i-1], levels[i], isIncreasing) {
			return false
		}
	}
	return true
}

func calcIncreasing(levels []int) bool {
	ups, downs := 0, 0
	for i := 1; i < len(levels); i++ {
		if checkInRange(levels[i-1], levels[i]) {
			if levels[i] > levels[i-1] {
				ups++
			} else if levels[i-1] > levels[i] {
				downs++
			}
		}
	}
	return ups > downs
}

func checkReportDampener(levels []int) bool {
	if checkReport(levels) {
		return true
	}
	for i := range levels {
		c := make([]int, len(levels))
		copy(c, levels)
		if checkReport(slices.Delete(c, i, i+1)) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	safeReports := 0
	safeDampenedReports := 0
	for scanner.Scan() {
		var levels []int
		levelsStr := strings.Fields(scanner.Text())
		for _, level := range levelsStr {
			num, _ := strconv.Atoi(level)
			levels = append(levels, num)
		}

		if checkReport(levels) {
			safeReports++
		}
		if checkReportDampener(levels) {
			safeDampenedReports++
		}
	}
	fmt.Println("part 1 reports:", safeReports)
	fmt.Println("part 2 reports:", safeDampenedReports)
}
