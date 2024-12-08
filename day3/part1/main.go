package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var mulRegexp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	total := 0

	for _, match := range mulRegexp.FindAllSubmatch(input, -1) {
		a, err := strconv.Atoi(string(match[1]))
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(string(match[2]))
		if err != nil {
			panic(err)
		}
		total += a * b
	}

	log.Printf("total: %d", total)
}
