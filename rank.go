package main

type Rank int

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const NumRanks = 13

var rankStrings = []string{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

func (r Rank) String() string {
	return rankStrings[r]
}

var rankShortStrings = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func (r Rank) ShortString() string {
	return rankShortStrings[r]
}

func (r Rank) Score() int {
	switch r {
	case Jack, Queen, King:
		return 10
	default:
		return int(r) + 1
	}
}
