package main

type Rank int

func (r Rank) Score() int {
	switch r {
	case Jack, Queen, King:
		return 10
	default:
		return int(r) + 1
	}
}

func (r Rank) ShortString() string {
	return ranks[r].shortString
}

func (r Rank) String() string {
	return ranks[r].string
}

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
	numRanks

	NumRanks = int(numRanks)
)

type rank struct {
	shortString, string string
}

var ranks = [numRanks]rank{
	{"A", "Ace"},
	{"2", "Two"},
	{"3", "Three"},
	{"4", "Four"},
	{"5", "Five"},
	{"6", "Six"},
	{"7", "Seven"},
	{"8", "Eight"},
	{"9", "Nine"},
	{"10", "Ten"},
	{"J", "Jack"},
	{"Q", "Queen"},
	{"K", "King"},
}
