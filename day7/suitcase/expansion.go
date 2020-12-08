package suitcase

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BagType string

type BagSet struct {
	BagType BagType
	Count   int
}

type ExpansionRule struct {
	ExpansionKey    BagType
	ExpansionValues []BagSet
}

var contentsRegex = regexp.MustCompile("([0-9]+) ([a-zA-Z\\s]+) bags?.?")

func (r *ExpansionRule) ParseRule(rule string) error {
	parts := strings.Split(rule, " bags contain ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid rule format: split on 'bags contain' but got %d != 2 parts", len(parts))
	}

	r.ExpansionKey = BagType(parts[0])

	if parts[1] == "no other bags." {
		r.ExpansionValues = nil
		return nil
	}

	contentsParts := strings.Split(parts[1], ", ")
	for _, part := range contentsParts {
		matches := contentsRegex.FindStringSubmatch(part)
		if len(matches) != 3 {
			return fmt.Errorf("invalid rule format: expected bag contents regex to have 3 matches, but got %d", len(matches))
		}
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("could not parse number in expansion rule (%v)", err)
		}

		r.ExpansionValues = append(r.ExpansionValues, BagSet{
			BagType: BagType(matches[2]),
			Count:   num,
		})
	}

	return nil
}
