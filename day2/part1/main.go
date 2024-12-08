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

	safeReports := 0

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")

		if len(values) < 2 {
			safeReports += 1
			continue
		}

		report := make([]int, len(values))
		for i, v := range values {
			var err error
			report[i], err = strconv.Atoi(v)
			if err != nil {
				log.Fatalf("invalid number: %s", v)
			}
		}

		checkFn := checkDec

		if report[0] < report[1] {
			checkFn = checkInc
		}

		safe := true
		for i := 0; i < len(report)-1 && safe; i++ {
			safe = checkFn(report[i], report[i+1])
		}
		if safe {
			safeReports += 1
			log.Printf("report: %v, safe: %v", report, safe)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	log.Printf("safeReports: %d", safeReports)
}
