package navigation

import (
	"strconv"
)

type heading int

const (
	East heading = iota
	North
	West
	South
)

//    E N W S
//    0 1 2 3
// x: 1 0 -1 0
// y: 0 1 0 -1

type ferryInstructionType string

const (
	FITMoveNorth   ferryInstructionType = "N"
	FITMoveSouth                        = "S"
	FITMoveEast                         = "E"
	FITMoveWest                         = "W"
	FITMoveForward                      = "F"
	FITRotateRight                      = "R"
	FITRotateLeft                       = "L"
)

type FerryInstruction struct {
	instructionType ferryInstructionType
	argument        int
}

func (fi *FerryInstruction) Parse(s string) error {
	fi.instructionType = ferryInstructionType(s[:1])
	arg, err := strconv.Atoi(s[1:])
	fi.argument = arg
	return err
}

type FerryLocation struct {
	x int
	y int
	w waypoint
}

func NewFerryLocation() *FerryLocation {
	return &FerryLocation{
		x: 0,
		y: 0,
		w: waypoint{
			10,
			1,
		},
	}
}

type waypoint struct {
	x int
	y int
}

func (fl *FerryLocation) ApplyInstruction(instruction FerryInstruction) {
	switch instruction.instructionType {
	case FITMoveForward:
		fl.x += fl.w.x * instruction.argument
		fl.y += fl.w.y * instruction.argument
	case FITMoveNorth:
		fl.w.y += instruction.argument
	case FITMoveSouth:
		fl.w.y -= instruction.argument
	case FITMoveEast:
		fl.w.x += instruction.argument
	case FITMoveWest:
		fl.w.x -= instruction.argument
	case FITRotateRight:
		instruction.argument = 360 - instruction.argument
		fallthrough
	case FITRotateLeft:
		for i := 0; i < instruction.argument; i += 90 {
			fl.w.x, fl.w.y = -1*fl.w.y, fl.w.x
		}
	}
}

func (fl *FerryLocation) DistanceFromOrigin() int {
	return abs(fl.x) + abs(fl.y)
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
