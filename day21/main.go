package main

import (
	"advent/common/lineprocessor"
	"advent/day21/allergen"
	"fmt"
	"strings"
)

func main() {
	p := allergen.NewProcessor()
	lineprocessor.ProcessLinesInFile("day21/input.txt", p)
	fmt.Printf("Confirmed non-allergen count: %d\n", p.ComputeNonAllergenCount())
	ingredients := p.ComputeDangerousIngredientsList()
	b := strings.Builder{}
	for _, a := range ingredients {
		b.WriteString(string(a))
		b.WriteRune(',')
	}
	l := b.String()
	l = l[:len(l)-1]
	fmt.Printf("Allergen list: %s\n", l)
}
