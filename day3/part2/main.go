package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var mulRegexp = regexp.MustCompile(`(?:do\(\))|(?:don't\(\))|(?:mul\((\d{1,3}),(\d{1,3})\))`)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	enable := true
	total := 0

	for _, match := range mulRegexp.FindAllSubmatch(input, -1) {
		if string(match[0]) == `do()` {
			enable = true
			continue
		}

		if string(match[0]) == `don't()` {
			enable = false
			continue
		}

		if !enable {
			continue
		}

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
