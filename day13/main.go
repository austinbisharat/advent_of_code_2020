package main

import (
	"advent/day13/bus"
	"advent/day13/modulo"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	curTime, schedules := parseFile("day13/input.txt")
	min, schedule := computeShortestWait(schedules, curTime)

	fmt.Printf("Wait until %d: %d min\n", schedule.ID, min)
	fmt.Printf("Product: %d\n", min*schedule.ID)

	t, err := computeEarliestTime(schedules)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Printf("Earliest time: %d\n", t)
	}
}

func computeShortestWait(schedules []bus.Schedule, curTime int) (int, bus.Schedule) {
	var min *int
	var schedule bus.Schedule
	for i := range schedules {
		s := schedules[i]
		t := s.TimeUntilNextBus(curTime)
		if min == nil || t < *min {
			min = &t
			schedule = s
		}
	}
	return *min, schedule
}

// Assumes all schedule ids are co-prime for now
func computeEarliestTime(schedules []bus.Schedule) (int, error) {
	if len(schedules) == 0 {
		return 0, nil
	}

	earliestTime := 0
	product := schedules[0].ID

	for _, schedule := range schedules[1:] {
		// product * x + earliestTime = schedule.Idx mod schedule.ID
		idx := schedule.Idx - schedule.ID
		remainder := modulo.Mod(0-idx-earliestTime, schedule.ID)
		x, gcd, err := modulo.SolveLinearCongruence(product, remainder, schedule.ID)
		if err != nil {
			return 0, err
		}

		earliestTime += product * x

		if (earliestTime+schedule.Idx)%schedule.ID != 0 {
			fmt.Printf("ERROR!!!")
		}
		product *= schedule.ID / gcd
	}

	return earliestTime, nil
}
func parseFile(filepath string) (int, []bus.Schedule) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	var schedules []bus.Schedule
	for i, schedule := range strings.Split(scanner.Text(), ",") {
		if schedule == "x" {
			continue
		}
		s, _ := strconv.Atoi(schedule)
		schedules = append(schedules, bus.Schedule{
			ID:  s,
			Idx: i,
		})
	}

	return num, schedules
}
