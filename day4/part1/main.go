package main

import (
	"bufio"
	"log"
	"os"
)

const XMAS = "XMAS"

func main() {
	board := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("reading standard input: %v", scanner.Err())
	}

	total := 0

	checkAny := func(i, j, incI, incJ int) bool {
		maxI := i + (len(XMAS)-1)*incI
		maxJ := j + (len(XMAS)-1)*incJ

		if maxI >= len(board) || maxI < 0 {
			return false
		}

		if maxJ >= len(board[i]) || maxJ < 0 {
			return false
		}

		for n := range XMAS {
			if XMAS[n] != board[i+n*incI][j+n*incJ] {
				return false
			}
		}

		return true
	}

	for i := range board {
		for j := range board[i] {
			// check South
			if checkAny(i, j, 1, 0) {
				total += 1
			}
			// check North
			if checkAny(i, j, -1, 0) {
				total += 1
			}
			// check East
			if checkAny(i, j, 0, 1) {
				total += 1
			}
			// check West
			if checkAny(i, j, 0, -1) {
				total += 1
			}
			// check South-East
			if checkAny(i, j, 1, 1) {
				total += 1
			}
			// check South-West
			if checkAny(i, j, 1, -1) {
				total += 1
			}
			// check North-East
			if checkAny(i, j, -1, 1) {
				total += 1
			}
			// check North-West
			if checkAny(i, j, -1, -1) {
				total += 1
			}
		}
	}

	log.Printf("Total: %d", total)
}
