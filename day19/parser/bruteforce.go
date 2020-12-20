package parser

import "strings"

type bruteForceParser struct {
	grammar *grammar

	validStrings map[string]bool
	expansions   [][]string
}

// can only handle non-cyclic grammars because it precomputes all
// possible strings
func NewBruteForceParser(file string) Parser {
	p := bruteForceParser{
		grammar: newGrammarFromFile(file),
	}
	p.init()
	return &p
}

func (p *bruteForceParser) init() {
	p.validStrings = make(map[string]bool)
	p.expansions = make([][]string, len(p.grammar.rules))

	t0 := grammarToken{
		isLiteral:  false,
		ruleNumber: 0,
	}
	for _, s := range p.expand(t0) {
		p.validStrings[s] = true
	}
}

func (p *bruteForceParser) expand(t grammarToken) []string {
	if t.isLiteral {
		return []string{t.text}
	}

	if p.expansions[t.ruleNumber] == nil {
		expansions := p.grammar.rules[t.ruleNumber].expansions
		expansionsForEachExpansion := make([][]string, 0, len(expansions))
		for _, e := range expansions {
			expanded := make([][]string, 0, len(e))
			for _, subToken := range e {
				expanded = append(expanded, p.expand(subToken))
			}
			var expandedText []string
			crossProduct(make([]string, len(expanded)), 0, expanded, &expandedText)
			expansionsForEachExpansion = append(expansionsForEachExpansion, expandedText)
		}

		p.expansions[t.ruleNumber] = flatten(expansionsForEachExpansion)
	}

	return p.expansions[t.ruleNumber]
}

func (p *bruteForceParser) IsValid(line string) bool {
	return p.validStrings[line]
}

func crossProduct(curProduct []string, i int, expansions [][]string, results *[]string) {
	if i == len(expansions) {
		b := strings.Builder{}
		for _, s := range curProduct {
			b.WriteString(s)
		}
		*results = append(*results, b.String())
	} else {
		for j := range expansions[i] {
			curProduct[i] = expansions[i][j]
			crossProduct(curProduct, i+1, expansions, results)
		}
	}
}

func flatten(arr [][]string) []string {
	var length int
	for _, subArr := range arr {
		length += len(subArr)
	}

	newArr := make([]string, 0, length)
	for _, subArr := range arr {
		newArr = append(newArr, subArr...)
	}
	return newArr
}
