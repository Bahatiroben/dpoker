package main

import (
	"fmt"
	"math/rand/v2"
)

type Card struct {
	suit   string // hearts, clabs, diamonds and spades
	symbol string // K,Q,J, A,7, 2,3,4,5,6
	value  int    // value count of the card
	color  string // red (hears, diamonds) or black (clabs, spades)
}

type Player struct {
	name string
	cards []*Card
	totalValue int
	totalNumberOfCards int

}

// a map of cards with name as key and Card as the value
type TAllCards map[string]*Card

type Game struct {
	players []*Player
	numberofPlayers int
	cardsPerPlayer int
	baseCards map[int]string
	allCards TAllCards
	cardsList [] *Card
}

func (g *Game) disburseCards() {
	g.shuffleDeck()

	g.cardsPerPlayer = len(g.cardsList) / g.numberofPlayers

	for index, card:= range g.cardsList {
		currentPlayerPosition := index % g.numberofPlayers
		if len(g.players) < g.numberofPlayers {
			g.players = append(g.players, &Player{ name: fmt.Sprintf("Player #%v", index + 1)})
		}
		g.players[currentPlayerPosition].cards = append(g.players[currentPlayerPosition].cards, card)
		g.players[currentPlayerPosition].totalValue += card.value
		g.players[currentPlayerPosition].totalNumberOfCards += 1
	}
}

func (g *Game) shuffleDeck() {
	rand.Shuffle(len(g.cardsList), func(i int, j int) {
		g.cardsList[i], g.cardsList[j] = g.cardsList[j], g.cardsList[i]
	})
}

// creates the cards from the original setup of card values
func (g *Game) loadAllCards() {
	g.allCards = TAllCards {}
	for value, symbol := range g.baseCards {

		// load the reds (hearts)
		g.allCards[symbol+" - Hearts"] = &Card{
			suit:   "Hearts",
			symbol: symbol,
			value:  value,
			color:  "Red",
		}

		// load the reds (diamonds)
		g.allCards[symbol+" - Diamonds"] = &Card{
			suit:   "Diamonds",
			symbol: symbol,
			value:  value,
			color:  "Red",
		}

		// load the blacks (clabs)
		g.allCards[symbol+" - Clabs"] = &Card{
			suit:   "Clabs",
			symbol: symbol,
			value:  value,
			color:  "Black",
		}

		// load the blacks (spades)
		g.allCards[symbol+" - Spades"] = &Card{
			suit:   "Spades",
			symbol: symbol,
			value:  value,
			color:  "Black",
		}
	}
}

func (g *Game) packTheCardsList() {
	g.cardsList = []*Card {}
	for _, val := range g.allCards {
		g.cardsList = append(g.cardsList, val)
	}
}

// numberOfPlayers 2,3,4,6
func NewGame(numberOfPlayers int) *Game {
	game := Game{
		numberofPlayers: numberOfPlayers,
		baseCards: map[int]string{
			2: "2", 3: "3", 4: "4", 5: "5", 6: "6",
			7: "7", 8: "Q", 9: "J", 10: "K", 11: "A",
		},
	}
	game.loadAllCards()
	game.packTheCardsList()
	return &game
}

func main() {

	game1 := NewGame(5)
	game1.shuffleDeck()
	game1.disburseCards()

	for _, player := range game1.players {
		fmt.Println(player.name, player.totalValue, player.totalNumberOfCards)
	}

}
