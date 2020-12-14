package main

import (
	"advent/common/lineprocessor"
	"advent/day14/bitmask"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	runtime := &ProgramRuntime{
		memory:   bitmask.NewProgramMemory(),
		memoryV2: bitmask.NewProgramMemory(),
	}

	lineprocessor.ProcessLinesInFile("day14/input.txt", runtime)
	fmt.Printf("Memory sum: %d\n", runtime.memory.GetSum())
	fmt.Printf("Memory v2 sum: %d\n", runtime.memoryV2.GetSum())
}

type ProgramRuntime struct {
	memory   *bitmask.ProgramMemory
	memoryV2 *bitmask.ProgramMemory
}

var maskRegex = regexp.MustCompile("^mask = ([X10]{36})$")
var memRegex = regexp.MustCompile("^mem\\[([0-9]+)] = ([0-9]+)$")

func (r *ProgramRuntime) ProcessLine(_ int, line string) error {

	if maskSubmatches := maskRegex.FindStringSubmatch(line); maskSubmatches != nil {
		var mask bitmask.Mask
		err := mask.Parse(maskSubmatches[1])
		if err != nil {
			return err
		}
		r.memory.UpdateMask(mask)
		r.memoryV2.UpdateMask(mask)
	} else if memSubmatches := memRegex.FindStringSubmatch(line); memSubmatches != nil {
		address, err := strconv.ParseUint(memSubmatches[1], 10, 64)
		if err != nil {
			return err
		}
		value, err := strconv.ParseUint(memSubmatches[2], 10, 64)
		if err != nil {
			return err
		}
		r.memory.UpdateMemory(address, value)
		r.memoryV2.UpdateMemoryV2(address, value)
	} else {
		return fmt.Errorf("invalid line: %s", line)
	}

	return nil
}
