package main

import (
	"blackjack/deck"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Hand - list of cards
type Hand []deck.Card

// String - returns hand cards string representation
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range strs {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString - returns dealer hand string representation while game is not over
func (h Hand) DealerString() string {
	return fmt.Sprintf("%s, **Hidden**", h[0].String())
}

// MinScore - minimum possible score
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += Min(int(c.Rank), 10)
	}
	return score
}

// Score - actual possible score
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, card := range h {
		if card.Rank == deck.Ace {
			// ace are currently == 1, and we are changing it to 11
			return minScore + 10
		}
	}
	return minScore
}

// State - current game state
type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// GameState - game state struct
type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

// CurrentPlayer - returns pointer to current player hand
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("there's no current player")
	}
}

// clone - cloning game state struct
func clone(gs GameState) GameState {
	res := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(res.Deck, gs.Deck)
	copy(res.Player, gs.Player)
	copy(res.Dealer, gs.Dealer)
	return res
}

// Shuffle - shuffles current game state deck
func Shuffle(gs GameState) GameState {
	res := clone(gs)
	res.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return res
}

// Deal - creating starting hands
func Deal(gs GameState) GameState {
	var card deck.Card
	res := clone(gs)
	res.Player = make(Hand, 0, 2)
	res.Dealer = make(Hand, 0, 2)
	for i := 0; i < 2; i++ {
		card, res.Deck = draw(res.Deck)
		res.Player = append(res.Player, card)
		card, res.Deck = draw(res.Deck)
		res.Dealer = append(res.Dealer, card)
	}
	res.State = StatePlayerTurn
	return res
}

// Hit - adding card to  hand
func Hit(gs GameState) GameState {
	var card deck.Card
	res := clone(gs)
	hand := res.CurrentPlayer()
	card, res.Deck = draw(res.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(res)
	}
	return res
}

// Stand - stop adding cards to a hand and change game state
func Stand(gs GameState) GameState {
	res := clone(gs)
	res.State++
	return res
}

// End - print winner and final score, also clear players and dealers hand
func End(gs GameState) GameState {
	res := clone(gs)
	pScore, dScore := res.Player.Score(), res.Dealer.Score()
	fmt.Println("***FINAL HANDS***")
	fmt.Println(fmt.Sprintf("Player: %s\nScore: %d", res.Player, pScore))
	fmt.Println(fmt.Sprintf("Player: %s\nScore: %d", res.Dealer, dScore))
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	res.Player = nil
	res.Dealer = nil
	return res
}

// draw - process adding card to a hand
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// Min - minimum int value
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Want to play a blackjack game y/n ?")
	for sc.Scan() {
		choice := sc.Text()
		switch strings.TrimSpace(strings.ToLower(choice)) {
		case "y":
			var (
				gs    GameState
				input string
			)

			gs = Shuffle(gs)
			gs = Deal(gs)

			for gs.State == StatePlayerTurn {
				fmt.Println("Player: ", gs.Player)
				fmt.Println("Dealer: ", gs.Dealer.DealerString())
				fmt.Println("What to do? Hit(h)/Stand(s)")
				fmt.Scanf("%s\n", &input)
				switch input {
				case "h":
					gs = Hit(gs)
				case "s":
					gs = Stand(gs)
				default:
					fmt.Println("Invalid option: ", input)
				}
			}

			for gs.State == StateDealerTurn {
				if (gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17)) && gs.Player.Score() < 21 {
					gs = Hit(gs)
				} else {
					gs = Stand(gs)
				}
			}

			gs = End(gs)
		case "n":
			return
		default:
			fmt.Println("Invalid option")
		}
		fmt.Println("New Game y/n ?")
	}

}
