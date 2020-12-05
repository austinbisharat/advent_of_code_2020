package main

type Ticket uint16

const (
	MaxPossibleTicket Ticket = 0b1111111111
	MinPossibleTicket Ticket = 0
)

func TicketFromStr(ticketStr string) Ticket {
	var ticket Ticket
	for i, char := range ticketStr {

		if char == 'B' || char == 'R' {
			ticket |= 1 << uint8(len(ticketStr) - i - 1)
		}
	}

	return ticket
}
