package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	entries := getEntries("day2/password_input.txt")

	count := 0
	for _, entry := range entries {
		if entry.policy.isValidPassword2(entry.password) {
			count++
		}
	}

	fmt.Printf("(%d)/(%d) passwords conformed to their policy", count, len(entries))

}

type PasswordPolicy struct {
	low int
	high int
	character rune
	charStr string
}

func (p * PasswordPolicy) isValid() bool {
	return p.low >= 0 && p.low <= p.high && len(p.charStr) == 1
}

func (p * PasswordPolicy) isValidPassword(password string) bool {
	var count int
	for _, c := range password {
		if c == p.character {
			count++
		}
	}

	return count >= p.low && count <= p.high
}

func (p * PasswordPolicy) isValidPassword2(password string) bool {
	return (password[p.low-1] == p.charStr[0]) != (password[p.high-1] == p.charStr[0])
}

type PasswordEntry struct {
	policy PasswordPolicy
	password string
}

func (p* PasswordEntry) parse(line string) error {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return errors.New("expected exactly 2 parts to line")
	}

	policyText := parts[0]
	passwordText := parts[1]

	policyParts := strings.Split(policyText, " ")

	if len(policyParts) != 2 {
		return errors.New("expected exactly 2 parts to policy text")
	}

	rangeText := policyParts[0]
	charText := policyParts[1]

	if len(charText) != 1 {
		return errors.New("expected exactly 1 char in policy text")
	}

	char, _ := utf8.DecodeRuneInString(charText)

	rangeParts := strings.Split(rangeText, "-")

	if len(rangeParts) != 2 {
		return errors.New("expected exactly 2 parts to range text")
	}

	low, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return err
	}

	high, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return err
	}

	p.policy = PasswordPolicy{
		low,
		high,
		char,
		charText,
	}

	if !p.policy.isValid() {
		return errors.New("invald policy")
	}

	p.password = strings.TrimSpace(passwordText)

	return nil
}

func getEntries(path string) []PasswordEntry {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries := make([]PasswordEntry, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entryText := scanner.Text()

		if len(entryText) == 0 {
			continue
		}

		var entry PasswordEntry
		err = entry.parse(entryText)
		if err != nil {
			log.Fatal(err)
		}

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries
}

