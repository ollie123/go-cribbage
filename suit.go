package main

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

const NumSuits = 4

var suitStrings = []string{
	"Clubs",
	"Diamonds",
	"Hearts",
	"Spades",
}

var suitColors = []Color{
	BrightWhite,
	Red,
	Red,
	BrightWhite,
}

func (s Suit) Color() Color {
	return suitColors[s]
}

func (s Suit) String() string {
	return suitStrings[s]
}

var suitShortStrings = []string{"♣", "♦", "♥", "♠"}

func (s Suit) ShortString() string {
	return suitShortStrings[s]
}
