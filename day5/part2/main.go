package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	X string
	Y string
}

func (r Rule) IsSatisfied(pages []string) (bool, int, int) {
	xPos := slices.Index(pages, r.X)
	yPos := slices.Index(pages, r.Y)
	if xPos == -1 || yPos == -1 {
		return true, xPos, yPos
	}
	return xPos < yPos, xPos, yPos
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	rules := []Rule{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.SplitN(scanner.Text(), "|", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid rule: %s", scanner.Text())
		}
		rules = append(rules, Rule{X: parts[0], Y: parts[1]})
	}

	if scanner.Err() != nil {
		log.Fatalf("reading standard input: %v", scanner.Err())
	}

	total := 0

	for scanner.Scan() {
		pages := strings.Split(scanner.Text(), ",")
		fixed := false

		for i := 0; i < len(rules); i++ {
			ok, posX, posY := rules[i].IsSatisfied(pages)
			if !ok {
				fixed = true
				tmp := pages[posX]
				pages[posX] = pages[posY]
				pages[posY] = tmp
				i = -1
			}
		}

		if !fixed {
			continue
		}
		mid, err := strconv.Atoi(pages[len(pages)/2])
		if err != nil {
			log.Fatalf("Invalid mid page number: %s", scanner.Text())
		}
		total += mid
	}

	if scanner.Err() != nil {
		log.Fatalf("reading standard input: %v", scanner.Err())
	}

	log.Printf("Total: %d", total)
}
