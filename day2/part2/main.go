package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkDec(a, b int) bool {
	return b < a && a-b <= 3
}

func checkInc(a, b int) bool {
	return a < b && b-a <= 3
}

func checkLevel(level []int) bool {
	checkFn := checkDec

	if level[0] < level[1] {
		checkFn = checkInc
	}

	for i := 0; i < len(level)-1; i++ {
		if !checkFn(level[i], level[i+1]) {
			return false
		}
	}
	return true
}

func fullCheckLevel(level []int) bool {
	if checkLevel(level) {
		return true
	}

	for i := range level {
		var subLevel []int
		if i > 0 {
			subLevel = append(subLevel, level[:i]...)
		}
		if i < len(level)-1 {
			subLevel = append(subLevel, level[i+1:]...)
		}
		if checkLevel(subLevel) {
			return true
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	safeLevels := 0

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")

		if len(values) < 2 {
			safeLevels += 1
			continue
		}

		level := make([]int, len(values))
		for i, v := range values {
			var err error
			level[i], err = strconv.Atoi(v)
			if err != nil {
				log.Fatalf("invalid number: %s", v)
			}
		}

		if fullCheckLevel(level) {
			safeLevels += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	log.Printf("safeLevels: %d", safeLevels)
}
