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

func checkReport(report []int) bool {
	checkFn := checkDec

	if report[0] < report[1] {
		checkFn = checkInc
	}

	for i := 0; i < len(report)-1; i++ {
		if !checkFn(report[i], report[i+1]) {
			return false
		}
	}
	return true
}

func fullCheckReport(report []int) bool {
	if checkReport(report) {
		return true
	}

	for i := range report {
		var subReport []int
		if i > 0 {
			subReport = append(subReport, report[:i]...)
		}
		if i < len(report)-1 {
			subReport = append(subReport, report[i+1:]...)
		}
		if checkReport(subReport) {
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

		report := make([]int, len(values))
		for i, v := range values {
			var err error
			report[i], err = strconv.Atoi(v)
			if err != nil {
				log.Fatalf("invalid number: %s", v)
			}
		}

		if fullCheckReport(report) {
			safeLevels += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	log.Printf("safeLevels: %d", safeLevels)
}
