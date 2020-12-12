package seatingchart

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type seat uint16

const (
	seatEmpty    = 0b00000
	seatOccupied = 0b00001
	seatFloor    = 0b10000
)

type location struct {
	row int
	col int
}

type locationSet []location

type SeatingChart struct {
	seats               [][]seat
	neighborsByLocation [][]locationSet
}

type direction uint8

const (
	right     direction = 0
	rightDown direction = 1
	down      direction = 2
	leftDown  direction = 3
)

func shiftByDistInDirection(og location, dist int, dir direction) location {
	if dir == right {
		return location{
			og.row,
			og.col + dist,
		}
	} else if dir == rightDown {
		return location{
			og.row + dist,
			og.col + dist,
		}
	} else if dir == down {
		return location{
			og.row + dist,
			og.col,
		}
	} else {
		return location{
			og.row + dist,
			og.col - dist,
		}
	}
}

func (sc *SeatingChart) locationIsValid(loc location) bool {
	return loc.row >= 0 && loc.row < len(sc.seats) && loc.col >= 0 && loc.col < len(sc.seats[loc.row])
}

func (sc *SeatingChart) computeNeighbors() {
	sc.neighborsByLocation = make([][]locationSet, len(sc.seats))
	for row := range sc.seats {
		sc.neighborsByLocation[row] = make([]locationSet, len(sc.seats[row]))
	}

	for row := range sc.seats {
		for col := range sc.seats[row] {
			if sc.seats[row][col] == seatFloor {
				continue
			}

			loc := location{
				row: row,
				col: col,
			}

			stopSearchDirection := [4]bool{}

			for dist := 1; !(stopSearchDirection[0] && stopSearchDirection[1] && stopSearchDirection[2] && stopSearchDirection[3]); dist++ {
				for dir := direction(0); dir <= leftDown; dir++ {
					if stopSearchDirection[dir] {
						continue
					}

					potentialNeighborLocation := shiftByDistInDirection(loc, dist, dir)
					if !sc.locationIsValid(potentialNeighborLocation) {
						stopSearchDirection[dir] = true
						continue
					}

					neighbor := sc.seats[potentialNeighborLocation.row][potentialNeighborLocation.col]
					if neighbor == seatEmpty || neighbor == seatOccupied {
						stopSearchDirection[dir] = true
						sc.neighborsByLocation[loc.row][loc.col] =
							append(sc.neighborsByLocation[loc.row][loc.col], potentialNeighborLocation)

						sc.neighborsByLocation[potentialNeighborLocation.row][potentialNeighborLocation.col] =
							append(sc.neighborsByLocation[potentialNeighborLocation.row][potentialNeighborLocation.col], loc)
					}
				}
			}
		}
	}
}

func (sc *SeatingChart) forEachNeighbor(row, col int, f seatFunc) {
	for _, n := range sc.neighborsByLocation[row][col] {
		f(n.row, n.col, sc.seats[n.row][n.col])
	}
}

func (sc *SeatingChart) forEachSeat(f seatFunc) {
	for r, row := range sc.seats {
		for c, seat := range row {
			f(r, c, seat)
		}
	}
}

func (sc *SeatingChart) copy() *SeatingChart {
	cpy := &SeatingChart{
		seats:               make([][]seat, len(sc.seats)),
		neighborsByLocation: sc.neighborsByLocation,
	}

	for i := range sc.seats {
		cpy.seats[i] = make([]seat, len(sc.seats[i]))
		copy(cpy.seats[i], sc.seats[i])
	}
	return cpy
}

type seatFunc func(int, int, seat)

type FerrySimulator struct {
	seatingCharts [2]*SeatingChart
	curStep       int
}

func NewFerryChartSimulator(path string) *FerrySimulator {
	initialSeatingChart := readSeatingChartFromFile(path)
	initialSeatingChart.computeNeighbors()
	return &FerrySimulator{
		seatingCharts: [2]*SeatingChart{
			initialSeatingChart,
			initialSeatingChart.copy(),
		},
		curStep: 0,
	}
}

func readSeatingChartFromFile(path string) *SeatingChart {
	sc := &SeatingChart{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		sc.seats = append(sc.seats, make([]seat, len(line)))
		for j, s := range line {
			switch s {
			case 'L':
				sc.seats[i][j] = seatEmpty
			case '#':
				sc.seats[i][j] = seatOccupied
			case '.':
				sc.seats[i][j] = seatFloor
			default:
				log.Fatalf("Unrecognized seat type: %d", s)
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sc
}

func (s *FerrySimulator) Step() bool {
	changed := false
	curSeatingChart := s.seatingCharts[s.curStep%2]
	nextSeatingChart := s.seatingCharts[(s.curStep+1)%2]
	curSeatingChart.forEachSeat(func(row int, col int, curSeat seat) {
		if curSeat == seatFloor {
			return
		}

		var sum uint16
		curSeatingChart.forEachNeighbor(row, col, func(a int, b int, neighbor seat) {
			sum += uint16(neighbor) & 1
		})

		if sum == 0 && curSeat == seatEmpty {
			nextSeatingChart.seats[row][col] = seatOccupied
			changed = true
		} else if sum >= 5 && curSeat == seatOccupied {
			nextSeatingChart.seats[row][col] = seatEmpty
			changed = true
		} else {
			nextSeatingChart.seats[row][col] = curSeat
		}
	})
	s.curStep++
	return changed
}

func (s *FerrySimulator) Run() {
	for s.Step() {
		fmt.Printf("On step %d\n", s.curStep)
	}
}

func (s *FerrySimulator) CountOccupiedSeats() uint {
	var total uint
	s.seatingCharts[s.curStep%2].forEachSeat(func(_ int, _ int, s seat) {
		total = total + (uint(s) & 0b1)
	})
	return total
}

func (s *FerrySimulator) PrintCurrentState() {
	seatMap := map[seat]rune{
		seatEmpty:    'L',
		seatOccupied: '#',
		seatFloor:    '.',
	}
	cur := s.seatingCharts[s.curStep%2].seats
	for r := range cur {
		for c := range cur[r] {
			fmt.Printf("%s", string(seatMap[cur[r][c]]))
		}
		fmt.Println()
	}
}
