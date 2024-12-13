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

type Coordinates struct {
	X int
	Y int
}

type Pos struct {
	Coordinates
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
	startCoordinates := pos.Coordinates

	checked := make([]Coordinates, 0)
	loopCount := 0

	for {
		log.Printf("Current position: %v", pos)
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

		// If the next position is not the starting position and has not already been checked
		// check if adding an obstacle in the next position would induce a loop
		if next.Coordinates != startCoordinates && !slices.Contains(checked, next.Coordinates) {
			checked = append(checked, next.Coordinates)
			if checkNextPosMakeLoop(duplicateLab(lab), *pos) {
				loopCount++
			}
		}

		// if the next position is empty, move to it
		pos = &next
	}

	log.Printf("Found %d loops", loopCount)
}

func checkNextPosMakeLoop(lab [][]byte, pos Pos) bool {
	log.Printf("Checking if adding an obstacle at %v would induce a loop", pos.GetNextPos())
	// add an obstacle in the next position
	next := pos.GetNextPos()
	lab[next.Y][next.X] = '#'

	history := make([]Pos, 0)

	for {
		if slices.Contains(history, pos) {
			return true
		}
		history = append(history, pos)
		// get the next position
		next := pos.GetNextPos()
		// if the next position is out of bounds, this is not a loop
		if next.Y < 0 || next.Y >= len(lab) || next.X < 0 || next.X >= len(lab[next.Y]) {
			return false
		}
		// if the next position is a wall, turn right
		if lab[next.Y][next.X] == '#' {
			pos.TurnRight()
			continue
		}
		// if the next position is empty, move to it
		pos = next
	}
}

func duplicateLab(ogLab [][]byte) [][]byte {
	lab := make([][]byte, len(ogLab))
	for i, row := range ogLab {
		lab[i] = slices.Clone(row)
	}
	return lab
}

func findInitPos(lab [][]byte) *Pos {
	for i, row := range lab {
		for j, cell := range row {
			if cell == '^' {
				return &Pos{
					Coordinates: Coordinates{X: j, Y: i},
					D:           North,
				}
			}
		}
	}
	return nil
}
