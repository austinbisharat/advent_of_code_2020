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
	l := formatList(ingredients)
	fmt.Printf("Allergen list: %s\n", l)
}

func formatList(ingredients []allergen.Ingredient) string {
	b := strings.Builder{}
	for _, a := range ingredients {
		b.WriteString(string(a))
		b.WriteRune(',')
	}
	l := b.String()
	return l[:len(l)-1]
}
