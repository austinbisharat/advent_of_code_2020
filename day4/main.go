package main

import (
	"advent/common/lineprocessor"
	"fmt"
)

func main() {

	processor := &PassportValidatorLineProcessor{}
	lineprocessor.ProcessLinesInFile("day4/input.txt", processor)
	fmt.Printf("Valid passport count: %v\n", processor.GetValidCount())

}

type PassportValidatorLineProcessor struct {
	currentPassportBuilder *PassportBuilder
	validPassportCount int
}

func (p *PassportValidatorLineProcessor) ProcessLine(_ int, line string) error {
	if len(line) > 0 {
		if p.currentPassportBuilder == nil {
			p.currentPassportBuilder = NewPassportBuilder()
		}

		p.currentPassportBuilder.withEntries(line)

	} else if p.currentPassportBuilder != nil {
		// at blank line and we have a passport from previous lines
		passport := p.currentPassportBuilder.build()
		if passport.isValid() {
			p.validPassportCount++
		}
		p.currentPassportBuilder = nil
	}

	return nil
}

func (p *PassportValidatorLineProcessor) GetValidCount() int {
	if p.currentPassportBuilder != nil {
		passport := p.currentPassportBuilder.build()
		if passport.isValid() {
			p.validPassportCount++
		}
	}
	return p.validPassportCount
}
