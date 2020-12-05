package main

import "fmt"

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
	for i := MinPossibleTicket; i < t.minTicket; i++ {
		t.allTickets ^= i
	}
	for i := MaxPossibleTicket; i > t.maxTicket; i-- {
		t.allTickets ^= i
	}
	return t.allTickets & MaxPossibleTicket
}
