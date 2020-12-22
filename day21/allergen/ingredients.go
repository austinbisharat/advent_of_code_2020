package allergen

import "strings"

type Ingredient string

type IngredientSet map[Ingredient]bool

func (set IngredientSet) Add(other IngredientSet) {
	for o := range other {
		set[o] = true
	}
}

func (set IngredientSet) Subtract(other IngredientSet) {
	for o := range other {
		delete(set, o)
	}
}

func (set IngredientSet) Intersect(other IngredientSet) {
	for ingredient := range set {
		if !other[ingredient] {
			delete(set, ingredient)
		}
	}
}

func NewIngredientSetFromList(list string) IngredientSet {
	ig := make(IngredientSet)
	for _, s := range strings.Split(list, " ") {
		ig[Ingredient(s)] = true
	}
	return ig
}
