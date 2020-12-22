package allergen

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type Allergen string

type processor struct {
	possibleAllergens map[Allergen]IngredientSet
	ingredientCount   map[Ingredient]int
}

func NewProcessor() *processor {
	return &processor{
		possibleAllergens: make(map[Allergen]IngredientSet),
		ingredientCount:   make(map[Ingredient]int),
	}
}

func (p *processor) ProcessLine(_ int, line string) error {
	parts := strings.Split(line, " (contains ")
	if len(parts) > 2 || len(parts) == 0 {
		return fmt.Errorf("invalid format")
	}

	ingredients := NewIngredientSetFromList(parts[0])
	var allergens []Allergen
	if len(parts) > 1 {
		allergenText := parts[1]
		allergenText = allergenText[:len(allergenText)-1]
		for _, a := range strings.Split(allergenText, ", ") {
			allergens = append(allergens, Allergen(a))
		}
	}

	p.processItem(ingredients, allergens)
	return nil
}

func (p *processor) processItem(ingredients IngredientSet, allergens []Allergen) {
	for ig := range ingredients {
		p.ingredientCount[ig]++
	}

	for _, allergen := range allergens {
		if allergenIngredients := p.possibleAllergens[allergen]; allergenIngredients != nil {
			allergenIngredients.Intersect(ingredients)
		} else {
			p.possibleAllergens[allergen] = make(IngredientSet)
			p.possibleAllergens[allergen].Add(ingredients)
		}
	}
}

func (p *processor) computeAllPossibleAllergens() IngredientSet {
	possibleAllergens := make(IngredientSet)
	for _, igs := range p.possibleAllergens {
		possibleAllergens.Add(igs)
	}
	return possibleAllergens
}

func (p *processor) ComputeNonAllergenCount() int {
	var count int
	possibleAllergens := p.computeAllPossibleAllergens()
	for ingredient, ingredientCount := range p.ingredientCount {
		if !possibleAllergens[ingredient] {
			count += ingredientCount
		}
	}
	return count
}

func (p *processor) ComputeDangerousIngredientsList() []Ingredient {
	var allergens []Allergen
	for a := range p.possibleAllergens {
		allergens = append(allergens, a)
	}
	sort.Slice(allergens, func(i, j int) bool {
		return allergens[i] < allergens[j]
	})

	solution, hasSolution := p.computeIngredientsAssignment(nil, make(map[Ingredient]bool), allergens)
	if !hasSolution {
		log.Fatal("No solution")
	}
	return solution
}

func (p *processor) computeIngredientsAssignment(s []Ingredient, used map[Ingredient]bool, allergens []Allergen) ([]Ingredient, bool) {
	if len(s) == len(allergens) {
		return s, true
	}

	nextAllergen := allergens[len(s)]
	for ig := range p.possibleAllergens[nextAllergen] {
		if used[ig] {
			continue
		}

		s = append(s, ig)
		used[ig] = true

		solution, hasSolution := p.computeIngredientsAssignment(s, used, allergens)
		if hasSolution {
			return solution, true
		}

		delete(used, ig)
		s = s[:len(s)-1]
	}
	return nil, false
}
