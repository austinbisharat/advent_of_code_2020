package ticket

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Validator interface {
	IsValid(Ticket) (bool, int)
}

type rangeValidator struct {
	low  int
	high int
}

func (r *rangeValidator) fieldIsValid(field int) bool {
	return field >= r.low && field <= r.high
}

type FieldValidator struct {
	FieldName string
	ranges    []*rangeValidator
}

var fieldValidatorRegex = regexp.MustCompile("^([a-zA-Z ]*): ([0-9]*)-([0-9]*) or ([0-9]*)-([0-9]*)$")

func (fv *FieldValidator) Parse(s string) error {
	submatches := fieldValidatorRegex.FindStringSubmatch(s)
	if len(submatches) != 6 {
		return fmt.Errorf("invalid field validator string: '%s'", s)
	}
	fv.FieldName = submatches[1]

	low, err := strconv.Atoi(submatches[2])
	if err != nil {
		return fmt.Errorf("invalid field validator string: '%s'", s)
	}
	high, err := strconv.Atoi(submatches[3])
	if err != nil {
		return fmt.Errorf("invalid field validator string: '%s'", s)
	}

	fv.ranges = append(fv.ranges, &rangeValidator{
		low:  low,
		high: high,
	})

	low, err = strconv.Atoi(submatches[4])
	if err != nil {
		return fmt.Errorf("invalid field validator string: '%s'", s)
	}
	high, err = strconv.Atoi(submatches[5])
	if err != nil {
		return fmt.Errorf("invalid field validator string: '%s'", s)
	}

	fv.ranges = append(fv.ranges, &rangeValidator{
		low:  low,
		high: high,
	})
	return nil
}

func (fv *FieldValidator) IsFieldValid(field int) bool {
	for _, r := range fv.ranges {
		if r.fieldIsValid(field) {
			return true
		}
	}
	return false
}

type anyFieldTicketValidator struct {
	rangeValidators []*rangeValidator
}

func (a anyFieldTicketValidator) IsValid(ticket Ticket) (bool, int) {
	for _, field := range ticket {
		fieldIsValid := false
		for _, r := range a.rangeValidators {
			if r.fieldIsValid(field) {
				fieldIsValid = true
				break
			}
		}

		if !fieldIsValid {
			return false, field
		}
	}
	return true, 0
}

func NewAnyFieldTicketValidator(fieldValidators []FieldValidator) Validator {
	var rangeValidators []*rangeValidator

	for _, fieldValidator := range fieldValidators {
		for i := range fieldValidator.ranges {
			rv := *fieldValidator.ranges[i]
			rangeValidators = append(rangeValidators, &rv)
		}
	}

	sort.Slice(rangeValidators, func(i, j int) bool {
		return rangeValidators[i].low < rangeValidators[j].low
	})

	collapsedRangeValidators := []*rangeValidator{rangeValidators[0]}
	rangeValidators = rangeValidators[1:]
	for len(rangeValidators) > 0 {
		cur := rangeValidators[0]
		rangeValidators = rangeValidators[1:]

		last := collapsedRangeValidators[len(collapsedRangeValidators)-1]

		if last.high+1 >= cur.low && cur.high > last.high {
			last.high = cur.high
		} else if last.high+1 < cur.low {
			collapsedRangeValidators = append(collapsedRangeValidators, cur)
		}
	}

	return &anyFieldTicketValidator{
		rangeValidators: collapsedRangeValidators,
	}
}
