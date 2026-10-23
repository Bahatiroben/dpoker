package game

type Deck struct {
	cards []*Card
}

func(d *Deck) fill() {
	// load Faces
	d.loadFaces()
	// fill the Numbers
	d.loadNumbers()
	// load the jockers
	d.loadJokers()
}

func (d *Deck) loadFaces() {
	for ksuit, vcolor := range Suits {
		for face := range Faces {
			card := Card{
				Value: 0, 
				Suit:  ksuit,
				Color: vcolor,
				Face:  face,
				Type:  FacesStr,
			}
			d.cards = append(d.cards, &card)
		}
	}
}

func (d *Deck) loadNumbers() {
	for ksuit, vcolor := range Suits {
		for number := range Numbers {
			card := Card{
				Value: 0, 
				Suit:  ksuit,
				Color: vcolor,
				Number:  number,
				Type:  NumbersStr,
			}
			d.cards = append(d.cards, &card)
		}
	}
}

func (d *Deck) loadJokers() {
	d.cards = append(d.cards, []*Card{
	{
			Type: JokerStr,
			Color: Black,
		},
	{
			Type: JokerStr,
			Color: Red,
		},
	}...)
}