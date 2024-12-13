package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Pos struct {
	X int
	Y int
	D Direction
}

func (p *Pos) TurnRight() {
	p.D = (p.D + 1) % 4
}

func (p *Pos) GetNextPos() Pos {
	next := *p
	switch p.D {
	case North:
		next.Y--
	case East:
		next.X++
	case South:
		next.Y++
	case West:
		next.X--
	}
	return next
}

func main() {
	lab := make([][]byte, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lab = append(lab, slices.Clone(scanner.Bytes()))
	}

	if scanner.Err() != nil {
		log.Fatalf("reading standard input: %v", scanner.Err())
	}

	pos := findInitPos(lab)
	if pos == nil {
		log.Fatalf("No initial position found")
	}

	// number of cells explored
	explored := 0

	for {
		// mark current cell as explored if not already
		if lab[pos.Y][pos.X] != 'X' {
			lab[pos.Y][pos.X] = 'X'
			explored++
		}

		// get the next position
		next := pos.GetNextPos()
		// if the next position is out of bounds, we are done
		if next.Y < 0 || next.Y >= len(lab) || next.X < 0 || next.X >= len(lab[next.Y]) {
			break
		}

		// if the next position is a wall, turn right
		if lab[next.Y][next.X] == '#' {
			pos.TurnRight()
			continue
		}

		// if the next position is empty, move to it
		pos = &next
	}

	log.Printf("Explored: %d", explored)
}

func findInitPos(lab [][]byte) *Pos {
	for i, row := range lab {
		for j, cell := range row {
			if cell == '^' {
				return &Pos{X: j, Y: i, D: North}
			}
		}
	}
	return nil
}
