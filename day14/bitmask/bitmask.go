package bitmask

import "fmt"

type Mask struct {
	s     string
	ones  uint64
	zeros uint64
}

func (m *Mask) Parse(s string) error {
	if len(s) != 36 {
		return fmt.Errorf("invalid mask '%s'", s)
	}

	m.s = s

	m.ones = 0
	m.zeros = 0
	for i, char := range s {
		if char == '1' {
			m.ones |= 1 << (35 - i)
		} else if char == '0' {
			m.zeros |= 1 << (35 - i)
		}
	}

	m.zeros = ^m.zeros
	return nil
}

func (m *Mask) Apply(n uint64) uint64 {
	fmt.Printf("Applying:\n")
	fmt.Printf("\t %036b\n", n)
	fmt.Printf("\t %s\n", m.s)
	fmt.Printf("\t %036b\n", n&m.zeros|m.ones)
	n &= m.zeros
	n |= m.ones
	return n
}

type ProgramMemory struct {
	memory map[uint64]uint64
	mask   Mask

	curSum uint64
}

func NewProgramMemory() *ProgramMemory {
	return &ProgramMemory{
		memory: make(map[uint64]uint64),
		mask:   Mask{},
		curSum: 0,
	}
}

func (pm *ProgramMemory) UpdateMask(mask Mask) {
	pm.mask = mask
}

func (pm *ProgramMemory) UpdateMemory(address, value uint64) {
	fmt.Printf("Updating memory at address %d with %d\n", address, value)
	prevVal, exists := pm.memory[address]
	if exists {
		pm.curSum -= prevVal
	}
	value = pm.mask.Apply(value)
	pm.memory[address] = value
	pm.curSum += value
}

func (pm *ProgramMemory) GetSum() uint64 {
	var sum uint64
	for _, v := range pm.memory {
		sum += v
	}
	if pm.curSum != sum {
		fmt.Printf("ERRROR")
	}
	return pm.curSum
}
