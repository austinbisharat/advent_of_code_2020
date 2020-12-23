package main

import (
	"advent/day22/cards"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	player1Deck, player2Deck := ReadDecks("day22/input.txt")
	score := cards.PlayCombat(player1Deck, player2Deck)
	fmt.Printf("Winning score: %d\n", score)

	player1Deck, player2Deck = ReadDecks("day22/input.txt")
	score = cards.PlayRecursiveCombat(player1Deck, player2Deck)
	fmt.Printf("Winning score recursive: %d\n", score)
}

func ReadDecks(path string) (*cards.Deck, *cards.Deck) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	playerTexts := strings.Split(string(bytes), "\n\n")
	return readDeck(playerTexts[0]), readDeck(playerTexts[1])
}

func readDeck(text string) *cards.Deck {
	lines := strings.Split(text, "\n")
	lines = lines[1:]
	return cards.NewDeckFromText(lines)
}
