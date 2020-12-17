package main

import (
	"advent/common/lineprocessor"
	"advent/day16/ticket"
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	t := &TicketProcessor{}
	lineprocessor.ProcessLinesInFile("day16/input.txt", t)
	fieldValidators := t.fieldValidators
	myTicket, allTickets := t.myTicket, t.allTickets
	validTickets, errorRate := removeInvalidTickets(fieldValidators, allTickets)
	fmt.Printf("error rate %d\n", errorRate)

	fieldNames := computeFieldMapping(fieldValidators, validTickets)
	fmt.Printf("My ticket: %d\n", computeMyTicketValue(myTicket, fieldNames))
}

type TicketProcessor struct {
	phase int

	fieldValidators []ticket.FieldValidator
	myTicket        ticket.Ticket
	allTickets      []ticket.Ticket
}

func (t *TicketProcessor) ProcessLine(_ int, line string) error {
	if len(line) == 0 {
		return nil
	}

	if line == "your ticket:" {
		t.phase++
		return nil
	} else if line == "nearby tickets:" {
		t.phase++
		return nil
	}

	if t.phase == 0 {
		var fv ticket.FieldValidator
		err := fv.Parse(line)
		t.fieldValidators = append(t.fieldValidators, fv)
		return err
	}
	var ticket ticket.Ticket

	err := ticket.Parse(line)
	if err != nil {
		return err
	}

	if t.phase == 1 {
		t.myTicket = ticket
	}
	t.allTickets = append(t.allTickets, ticket)
	return err
}

func removeInvalidTickets(fieldValidators []ticket.FieldValidator, tickets []ticket.Ticket) (validTickets []ticket.Ticket, errorRate int) {
	v := ticket.NewAnyFieldTicketValidator(fieldValidators)
	for _, t := range tickets {
		isValid, invalidNum := v.IsValid(t)
		if isValid {
			validTickets = append(validTickets, t)
		} else {
			errorRate += invalidNum
		}
	}
	return
}

func computeFieldMapping(fieldValidators []ticket.FieldValidator, tickets []ticket.Ticket) []string {
	fieldIdxToValidFields := computeValidFieldsByFieldIndex(fieldValidators, tickets)

	sort.Slice(fieldIdxToValidFields, func(i, j int) bool {
		return len(fieldIdxToValidFields[i].possibleFieldNames) < len(fieldIdxToValidFields[j].possibleFieldNames)
	})

	fa := fieldAssignment{
		fieldAssignments: make([]string, len(tickets[0])),
		assignedFields:   map[string]bool{},
	}
	res, found := findFieldMapping(&fa, fieldIdxToValidFields)
	if !found {
		log.Fatal("!!! NOT FOUND !!!")
	}
	return res
}

func findFieldMapping(partialSolution *fieldAssignment, fieldIdxToValidFields []fieldIdxData) ([]string, bool) {
	if len(partialSolution.assignedFields) == len(fieldIdxToValidFields) {
		return partialSolution.fieldAssignments, true
	}

	nextFieldToAssign := fieldIdxToValidFields[len(partialSolution.assignedFields)]
	for possibleNextField := range nextFieldToAssign.possibleFieldNames {
		if partialSolution.assignedFields[possibleNextField] {
			continue
		}

		partialSolution.fieldAssignments[nextFieldToAssign.fieldIdx] = possibleNextField
		partialSolution.assignedFields[possibleNextField] = true
		res, found := findFieldMapping(partialSolution, fieldIdxToValidFields)
		if found {
			return res, found
		}
		delete(partialSolution.assignedFields, possibleNextField)
		partialSolution.fieldAssignments[nextFieldToAssign.fieldIdx] = ""

	}

	return nil, false
}

func computeValidFieldsByFieldIndex(fieldValidators []ticket.FieldValidator, tickets []ticket.Ticket) (fieldIdxToValidFields []fieldIdxData) {
	for fieldIdx := range tickets[0] {
		fieldNames := make(map[string]bool)
		for _, fv := range fieldValidators {
			fieldNames[fv.FieldName] = true
		}

		for _, ticket := range tickets {
			for _, fv := range fieldValidators {
				if !fv.IsFieldValid(ticket[fieldIdx]) {
					delete(fieldNames, fv.FieldName)
				}
			}
		}

		fieldIdxToValidFields = append(fieldIdxToValidFields, fieldIdxData{
			fieldIdx:           fieldIdx,
			possibleFieldNames: fieldNames,
		})
	}
	return fieldIdxToValidFields
}

func computeMyTicketValue(myTicket ticket.Ticket, fieldNames []string) int {
	product := 1
	for i, fieldName := range fieldNames {
		if strings.HasPrefix(fieldName, "departure") {
			product *= myTicket[i]
		}
	}
	return product
}

type fieldIdxData struct {
	fieldIdx           int
	possibleFieldNames map[string]bool
}

type fieldAssignment struct {
	fieldAssignments []string
	assignedFields   map[string]bool
}
