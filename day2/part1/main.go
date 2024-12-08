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

		checkFn := checkDec

		if level[0] < level[1] {
			checkFn = checkInc
		}

		safe := true
		for i := 0; i < len(level)-1 && safe; i++ {
			safe = checkFn(level[i], level[i+1])
		}
		if safe {
			safeLevels += 1
			log.Printf("level: %v, safe: %v", level, safe)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	log.Printf("safeLevels: %d", safeLevels)
}