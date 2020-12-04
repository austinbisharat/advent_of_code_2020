package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type PassportFieldKey string

const (
	PFKBirthYear      PassportFieldKey = "byr"
	PFKIssuedYear     PassportFieldKey = "iyr"
	PFKExpirationYear PassportFieldKey = "eyr"
	PFKHeight         PassportFieldKey = "hgt"
	PFKHairColor      PassportFieldKey = "hcl"
	PFKEyeColor       PassportFieldKey = "ecl"
	PFKPassportID     PassportFieldKey = "pid"
	PFKCountryID      PassportFieldKey = "cid"
)

var (
	hairColorRegex = regexp.MustCompile("^#[0-9a-f]{6}$")
	eyeColorRegex = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	passportIDRegex = regexp.MustCompile("^[0-9]{9}$")
)

var RequiredPassportFields = map[PassportFieldKey]PassportFieldValidator{
	PFKBirthYear: func(fieldValue string) bool {
		return isValidYear(fieldValue, 1920, 2002)
	},
	PFKIssuedYear: func(fieldValue string) bool {
		return isValidYear(fieldValue, 2010, 2020)
	},
	PFKExpirationYear: func(fieldValue string) bool {
		return isValidYear(fieldValue, 2020, 2030)
	},
	PFKHeight: func(fieldValue string) bool {
		if len(fieldValue) <= 2 {
			return false
		}
		numberText := fieldValue[0:len(fieldValue) - 2]
		suffix := fieldValue[len(fieldValue) - 2:]
		number, err := strconv.Atoi(numberText)
		if err != nil {
			return false
		}

		switch suffix {
		case "in":
			return number >= 59 && number <= 76
		case "cm":
			return number >= 150 && number <= 193
		default:
			return false
		}
	},
	PFKHairColor: func(fieldValue string) bool {
		return hairColorRegex.MatchString(fieldValue)
	},
	PFKEyeColor: func(fieldValue string) bool {
		return eyeColorRegex.MatchString(fieldValue)
	},
	PFKPassportID: func(fieldValue string) bool {
		return passportIDRegex.MatchString(fieldValue)
	},
}

func isValidYear(yearText string, minimum, maximum int) bool {
	year, err := strconv.Atoi(yearText)
	if err != nil {
		return false
	}

	return len(yearText) == 4 && year >= minimum && year <= maximum
}

type PassportFieldValidator func (fieldValue string) bool

type Passport map[PassportFieldKey]string

func (p *Passport) isValid() bool {
	for field, validator := range RequiredPassportFields {
		if fieldValue, exists := (*p)[field]; !exists || !validator(fieldValue) {
			return false
		}
	}

	return true
}

type PassportBuilder struct {
	passport Passport
}

func NewPassportBuilder() *PassportBuilder {
	return &PassportBuilder{
		passport: make(Passport),
	}
}

func (p *PassportBuilder) withEntries(entries string) {
	for _, entry := range strings.Fields(entries) {
		p.withEntry(entry)
	}
}

func (p *PassportBuilder) withEntry(entry string) {
	parts := strings.SplitN(entry, ":", 2)

	if len(parts) != 2 {
		log.Fatal("expected exactly two parts to entry")
	}

	p.passport[PassportFieldKey(parts[0])] = parts[1]
}

func (p *PassportBuilder) build() Passport {
	return p.passport
}

