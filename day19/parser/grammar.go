package parser

import (
	"advent/common/lineprocessor"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type grammar struct {
	rules []grammarRule
}

type grammarRule struct {
	ruleNumber int
	expansions []expansion
}

func (g *grammar) getRule(ruleNum int) grammarRule {
	return g.rules[ruleNum]
}

type expansion []grammarToken

type grammarToken struct {
	isLiteral  bool
	ruleNumber int
	text       string
}

func (r *grammarRule) parse(s string) error {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return fmt.Errorf("expected exactly 1 colon in rule")
	}

	var err error
	r.ruleNumber, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	for _, expansionText := range strings.Split(parts[1], "|") {
		expansionText = strings.TrimSpace(expansionText)
		tokenTexts := strings.Split(expansionText, " ")

		tokens := make([]grammarToken, len(tokenTexts))
		for i, t := range tokenTexts {
			if t[0] == '"' {
				tokens[i].isLiteral = true
				tokens[i].text = t[1 : len(t)-1]
			} else {
				tokens[i].ruleNumber, err = strconv.Atoi(t)
				if err != nil {
					return err
				}
			}
		}
		r.expansions = append(r.expansions, tokens)
	}

	return nil
}

type grammarReader struct {
	rules []grammarRule
}

func (reader *grammarReader) ProcessLine(_ int, line string) error {
	var rule grammarRule
	err := rule.parse(line)
	reader.rules = append(reader.rules, rule)
	return err
}

func newGrammarFromFile(file string) *grammar {
	g := grammarReader{}
	lineprocessor.ProcessLinesInFile(file, &g)
	rules := g.rules
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].ruleNumber < rules[j].ruleNumber
	})

	for i, rule := range rules {
		if i != rule.ruleNumber {
			log.Fatalf("Missing rule number %d\n", i)
		}
	}

	return &grammar{
		rules: rules,
	}
}
