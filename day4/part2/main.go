package main

import (
	"bufio"
	"log"
	"os"
)

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

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if board[i][j] != 'A' {
				continue
			}
			w := string([]byte{
				board[i-1][j-1],
				board[i-1][j+1],
				board[i+1][j-1],
				board[i+1][j+1],
			})
			switch w {
			case "MSMS", "MMSS", "SMSM", "SSMM":
				total++
			}
		}
	}

	log.Printf("Total: %d", total)
}
