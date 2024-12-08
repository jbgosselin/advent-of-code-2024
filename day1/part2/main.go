package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var rLine = regexp.MustCompile(`(\d+)\s+(\d+)`)

func main() {
	colA := make([]int, 0)
	colB := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		values := rLine.FindStringSubmatch(scanner.Text())
		if len(values) != 3 {
			log.Printf("invalid line: %s", scanner.Text())
			continue
		}
		a, err := strconv.Atoi(values[1])
		if err != nil {
			log.Printf("invalid number a: %s", values[1])
			continue
		}
		b, err := strconv.Atoi(values[2])
		if err != nil {
			log.Printf("invalid number b: %s", values[2])
			continue
		}
		colA = append(colA, a)
		colB = append(colB, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	if len(colA) != len(colB) {
		log.Fatalf("columns have different sizes: %d != %d", len(colA), len(colB))
	}

	slices.Sort(colA)
	slices.Sort(colB)

	total := 0

	for _, a := range colA {
		for _, b := range colB {
			if b == a {
				total += a
			}
		}
	}

	log.Printf("total: %d", total)
}
