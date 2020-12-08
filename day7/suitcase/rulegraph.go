package suitcase

import "advent/common/lineprocessor"

type RuleGraph interface {
	GetPossibleBagsContaining(bagType BagType) map[BagType]bool
	GetNumBagsIn(bagType BagType) int
}

func NewRuleGraph(filepath string) RuleGraph {
	g := &ruleGraph{
		expansionRuleByBagType: make(map[BagType]ExpansionRule),
		parentBagsByBagType:    make(map[BagType][]BagSet),

		numBagsIn: make(map[BagType]int),
	}

	lineprocessor.ProcessLinesInFile(filepath, g)
	return g
}

type ruleGraph struct {
	expansionRuleByBagType map[BagType]ExpansionRule
	parentBagsByBagType    map[BagType][]BagSet

	numBagsIn map[BagType]int
}

func (r *ruleGraph) ProcessLine(_ int, line string) error {

	var rule ExpansionRule
	err := rule.ParseRule(line)
	if err != nil {
		return err
	}

	r.addRule(rule)
	return nil
}

func (r *ruleGraph) addRule(rule ExpansionRule) {
	r.expansionRuleByBagType[rule.ExpansionKey] = rule

	for _, contents := range rule.ExpansionValues {
		r.parentBagsByBagType[contents.BagType] = append(r.parentBagsByBagType[contents.BagType], BagSet{
			BagType: rule.ExpansionKey,
			Count:   contents.Count,
		})
	}
}

func (r *ruleGraph) GetPossibleBagsContaining(bagType BagType) map[BagType]bool {
	toVisit := []BagType{bagType}
	discovered := make(map[BagType]bool)
	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = toVisit[1:]
		for _, parent := range r.parentBagsByBagType[cur] {
			if discovered[parent.BagType] {
				continue
			}
			discovered[parent.BagType] = true
			toVisit = append(toVisit, parent.BagType)
		}
	}

	return discovered
}

func (r *ruleGraph) GetNumBagsIn(bagType BagType) int {
	if numBags, exists := r.numBagsIn[bagType]; exists {
		return numBags
	} else {
		var sum int
		for _, child := range r.expansionRuleByBagType[bagType].ExpansionValues {
			sum += child.Count * (r.GetNumBagsIn(child.BagType) + 1)
		}
		r.numBagsIn[bagType] = sum
		return sum
	}
}
