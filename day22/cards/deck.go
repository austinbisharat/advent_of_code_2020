package cards

import (
	"strconv"
	"strings"
)

type Card uint8

type Deck struct {
	cards []Card
	start uint
	size  uint
}

func NewDeckWithCapacity(capacity int) *Deck {
	return &Deck{
		cards: make([]Card, capacity),
		start: 0,
		size:  0,
	}
}

func NewDeckFromText(text []string) *Deck {
	deck := NewDeckWithCapacity(len(text) * 2)
	for _, t := range text {
		cardNum, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		deck.InsertBottom(Card(cardNum))
	}
	return deck
}

func (d *Deck) Get(index uint) Card {
	index = (d.start + index) % uint(len(d.cards))
	return d.cards[index]
}

func (d *Deck) PopTop() Card {
	if d.size == 0 {
		panic("cannot pop empty deck")
	}
	d.size--
	idx := d.start
	d.start = (d.start + 1) % uint(len(d.cards))
	return d.cards[idx]
}

func (d *Deck) InsertBottom(card Card) {
	if d.size >= uint(len(d.cards)) {
		panic("deck too big!")
	}
	idx := (d.start + d.size) % uint(len(d.cards))
	d.size++
	d.cards[idx] = card
}

func (d *Deck) Len() uint {
	return d.size
}

func (d *Deck) Copy(num, capacity int) *Deck {
	newDeck := NewDeckWithCapacity(capacity)
	for i := 0; i < num; i++ {
		newDeck.InsertBottom(d.Get(uint(i)))
	}
	return newDeck
}

func (d *Deck) Key() string {
	b := strings.Builder{}
	for i := uint(0); i < d.size; i++ {
		b.WriteString(strconv.Itoa(int(d.Get(i))))
		b.WriteRune(',')
	}
	return b.String()
}
