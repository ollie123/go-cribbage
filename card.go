package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

func (c Card) ShortString() string {
	return c.Rank.ShortString() + c.Suit.ShortString()
}

func (c Card) ColorString() string {
	return c.Suit.Color().Format(c.Rank.ShortString() + c.Suit.ShortString())
}

func Random() Card {
	return Card{Rank(rand.Intn(NumRanks)), Suit(rand.Intn(NumSuits))}
}
