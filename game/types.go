package game

import "fmt"

type Card struct {
	Value int
	Suit  string // spades, hearts, Diamonds, clubs
	Color string

	Number int    // 2 through 10 // numbers
	Face   string // J, Q, K, A  faces

	Joker string // can only have a color, no number, or suit.
	Type  string
}

const Black = "Black"
const Red = "Red"

// key is a suit and value is a color
var Suits = map[string]string{
	"Clubs":    Black,
	"Spades":   Black,
	"Hearts":   Red,
	"Diamonds": Red,
}

var Numbers = map[int]int{
	2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10,
}

var Faces = map[string]string{
	"Jack":  "Jack",
	"Queen": "Queen",
	"King":  "King",
	"Ace":   "Ace",
}

var Types = map[string]string{
	"Faces":   "Faces",
	"Numbers": "Numbers",
	"Joker":   "Joker",
}

var Jokers = "Joker"
var Colors = map[string]string{Black: Black, Red: Red}

func LoadFaces() []*Card {
	faces := []*Card{}
	for ksuit, vcolor := range Suits {
		for face := range Faces {
			card := Card{
				Value: -1,
				Suit:  ksuit,
				Color: vcolor,
				Face:  face,
				Type:  "Faces",
			}
			fmt.Printf("%s of %s in %s\n", card.Face, card.Suit, card.Color)
			faces = append(faces, &card)
		}
	}
	return faces
}
