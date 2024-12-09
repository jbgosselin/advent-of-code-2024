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

	pos := &Pos{X: 0, Y: 0, D: North}

InitPos:
	for i, row := range lab {
		for j, cell := range row {
			if cell == '^' {
				pos.X = j
				pos.Y = i
				break InitPos
			}
		}
	}

	explored := 0

	for {
		if lab[pos.Y][pos.X] != 'X' {
			lab[pos.Y][pos.X] = 'X'
			explored++
		}

		next := pos.GetNextPos()
		if next.Y < 0 || next.Y >= len(lab) || next.X < 0 || next.X >= len(lab[next.Y]) {
			break
		}

		if lab[next.Y][next.X] == '#' {
			pos.TurnRight()
			continue
		}

		pos = &next
	}

	log.Printf("Explored: %d", explored)

}
