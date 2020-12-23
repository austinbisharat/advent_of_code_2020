package cards

import (
	"strings"
)

// PlayCombat plays a game of combat with the two decks
// and returns the winning players score
func PlayCombat(player1Deck, player2Deck *Deck) uint {
	for player1Deck.Len() > 0 && player2Deck.Len() > 0 {
		c1, c2 := player1Deck.PopTop(), player2Deck.PopTop()
		if c1 > c2 {
			player1Deck.InsertBottom(c1)
			player1Deck.InsertBottom(c2)
		} else {
			player2Deck.InsertBottom(c2)
			player2Deck.InsertBottom(c1)
		}
	}

	winner := player1Deck
	if player2Deck.Len() > player1Deck.Len() {
		winner = player2Deck
	}
	return computeScore(winner)
}

func PlayRecursiveCombat(player1Deck, player2Deck *Deck) uint {
	player1Won := playRecursiveCombat(player1Deck, player2Deck)
	winner := player1Deck
	if !player1Won {
		winner = player2Deck
	}
	return computeScore(winner)
}

func playRecursiveCombat(player1Deck, player2Deck *Deck) bool {
	gamesSeen := make(map[string]bool)
	for player1Deck.Len() > 0 && player2Deck.Len() > 0 {
		key := gameKey(player1Deck, player2Deck)
		if gamesSeen[key] {
			return true
		}
		gamesSeen[key] = true

		c1, c2 := player1Deck.PopTop(), player2Deck.PopTop()
		if player1Deck.Len() >= uint(c1) && player2Deck.Len() >= uint(c2) {
			newP1 := player1Deck.Copy(int(c1), int(c1+c2))
			newP2 := player2Deck.Copy(int(c2), int(c1+c2))
			if playRecursiveCombat(newP1, newP2) {
				player1Deck.InsertBottom(c1)
				player1Deck.InsertBottom(c2)
			} else {
				player2Deck.InsertBottom(c2)
				player2Deck.InsertBottom(c1)
			}
		} else if c1 > c2 {
			player1Deck.InsertBottom(c1)
			player1Deck.InsertBottom(c2)
		} else {
			player2Deck.InsertBottom(c2)
			player2Deck.InsertBottom(c1)
		}
	}
	return player1Deck.Len() > player2Deck.Len()
}

func gameKey(player1Deck, player2Deck *Deck) string {
	b := strings.Builder{}
	b.WriteString(player1Deck.Key())
	b.WriteString("::")
	b.WriteString(player2Deck.Key())
	return b.String()
}

func computeScore(winner *Deck) uint {

	var score uint
	for winner.Len() > 0 {
		score += winner.Len() * uint(winner.PopTop())
	}
	return score
}
