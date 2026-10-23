package game

import "fmt"

type Card struct {
	Value int
	Suit string // spades, hearts, Diamonds, clubs
	Color string

	Number int // 2 through 10 // numbers
	Face string // J, Q, K, A  faces

	Joker string // can only have a color, no number, or suit.
	Type string
}

var Suits = map[string]string{
	"Spades": "Spades",
	"Hearts": "Hearts",
	"Diamonds": "Diamonds",
	"Clubs": "Clubs",
}


var Numbers = map[int]int{
	2:2,3:3,4:4,5:5,6:6,7:7,8:8,9:9,10:10,
}

var Faces = map[string]string{
	"Jack": "Jack",
	"Queen": "Queen",
	"King": "King",
	"Ace": "Ace",
}

var Types = map[string]string {
	"Faces": "Faces",
	"Numbers": "Numbers",
	"Joker": "Joker",
}


var Jokers = "Joker"
var Colors = map[string]string{"Red": "Red", "Black": "Red"}

func LoadFaces() []*Card {
	faces := []*Card{}
	for color, _ := range Colors {
		for suit, _ := range Suits {
			for face, _ := range Faces {
				card := Card{
					Value: -1,
					Suit: suit,
					Color: color,
					Face: face,
					Type: "Faces",
				}
				fmt.Printf("%+v", card)
				faces = append(faces, &card)
			}
		}
	}
	return faces
}