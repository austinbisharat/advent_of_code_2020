package main

import (
	"fmt"
	"log"
)

func main() {
	processor := NewTicketLineProcessor()
	processLinesInFile("day5/input.txt", processor)
	fmt.Printf("Max ticket: %d\n", processor.MaxTicketSeen())
	fmt.Printf("Missing ticket: %d\n", processor.MissingTicket())
}

type ticketLineProcessor struct {
	maxTicket Ticket
	minTicket Ticket
	allTickets Ticket
}

func NewTicketLineProcessor() *ticketLineProcessor {
	return &ticketLineProcessor{
		MinPossibleTicket,
		MaxPossibleTicket,
		MinPossibleTicket,
	}
}

func (t *ticketLineProcessor) ProcessLine(_ int, line string) error {
	ticket := TicketFromStr(line)

	t.allTickets ^= ticket

	if ticket > t.maxTicket {
		t.maxTicket = ticket
	}

	if ticket < t.minTicket {
		t.minTicket = ticket
	}

	return nil
}

func (t *ticketLineProcessor) MaxTicketSeen() Ticket {
	return t.maxTicket
}

func (t *ticketLineProcessor) MissingTicket() Ticket {
	return t.allTickets ^ xorUpTo(t.minTicket) ^ t.minTicket ^ xorUpTo(t.maxTicket)
}

// returns result of doing xor of 0-n inclusive.
func xorUpTo(n Ticket) Ticket {
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n+1
	case 3:
		return 0
	default:
		log.Fatal("Should be impossible")
	}
	return 0
}
