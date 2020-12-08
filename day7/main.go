package main

import (
	"advent/day7/suitcase"
	"fmt"
)

const ShinyGoldBagType suitcase.BagType = "shiny gold"

func main() {
	g := suitcase.NewRuleGraph("day7/input.txt")

	fmt.Printf("%s bag type has %d possbible containing bags\n",
		ShinyGoldBagType,
		len(g.GetPossibleBagsContaining(ShinyGoldBagType)))

	fmt.Printf("%s bag type contains %d bags\n",
		ShinyGoldBagType,
		g.GetNumBagsIn(ShinyGoldBagType))
}
