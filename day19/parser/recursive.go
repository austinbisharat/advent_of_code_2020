package parser

type recursiveParser struct {
	grammar *grammar

	memo map[entry]bool
}

func NewRecursiveParser(file string) Parser {
	p := recursiveParser{
		grammar: newGrammarFromFile(file),
	}
	return &p
}

type entry struct {
	start int
	end   int
	rule  int
}

func (p *recursiveParser) IsValid(s string) bool {
	p.memo = make(map[entry]bool)
	t := []grammarToken{
		{
			isLiteral:  false,
			ruleNumber: 0,
			text:       "",
		},
	}
	return p.isValid(s, 0, len(s), t)
}

func (p *recursiveParser) isValid(s string, start, end int, tokens []grammarToken) bool {
	if len(tokens) == 0 {
		return len(s) == 0
	}

	firstToken := tokens[0]
	if len(tokens) == 1 {
		if firstToken.isLiteral {
			return firstToken.text == s[start:end]
		}

		ruleNum := firstToken.ruleNumber
		e := entry{
			start: start,
			end:   end,
			rule:  ruleNum,
		}

		ok, exists := p.memo[e]
		if exists {
			return ok
		}

		rule := p.grammar.getRule(ruleNum)
		for _, exp := range rule.expansions {
			ok = p.isValid(s, start, end, exp)
			if ok {
				p.memo[e] = true
				return true
			}
		}
		p.memo[e] = false
		return false
	}

	for i := start + 1; i < end; i++ {
		if p.isValid(s, start, i, tokens[:1]) && p.isValid(s, i, end, tokens[1:]) {
			return true
		}
	}

	return false
}
