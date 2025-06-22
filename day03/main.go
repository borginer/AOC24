package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	mulMaxSize    = len("mul(111,111)")
	enableMaxSize = len("don't()")
)

func mulParse(str string) (n int, length int) {
	str = strings.Split(str, ")")[0]

	if len(str) < len("mul(1,1") || str[0:4] != "mul(" {
		return
	}
	length = 4
	str = str[4:]

	nums := strings.Split(str, ",")
	if len(nums) != 2 {
		return
	}
	a, err1 := strconv.ParseInt(nums[0], 10, 64)
	b, err2 := strconv.ParseInt(nums[1], 10, 64)
	if err1 != nil || err2 != nil || a > 999 || b > 999 || a < 1 || b < 1 {
		return
	}

	return int(a * b), len(str) + 5
}

func checkEnabled(str string, cur bool) (enabled bool, length int) {
	str = strings.Split(str, ")")[0]
	if str == "do(" {
		return true, len("do()")
	} else if str == "don't(" {
		return false, len("don't()")
	}
	return cur, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	memory, err := reader.ReadString(0)
	if err != io.EOF {
		log.Fatal("oh boy")
	}

	mulSum, enabledMulSum, enabled := int64(0), int64(0), true
	for i := range memory {
		innerEnable, tokLen := checkEnabled(
			memory[i:min(i+enableMaxSize, len(memory))],
			enabled,
		)
		if tokLen > 0 {
			enabled = innerEnable
			i += (tokLen - 1)
			continue
		}

		n, l := mulParse(memory[i:min(i+mulMaxSize, len(memory))])
		i += (l - 1)
		mulSum += int64(n)
		if enabled {
			enabledMulSum += int64(n)
		}

	}

	fmt.Println("part 1: ", mulSum)
	fmt.Println("part 2: ", enabledMulSum)
}
