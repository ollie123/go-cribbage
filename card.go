package main

import (
	"fmt"
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (c Card) ColorString() string {
	return c.Suit.Color().Sprint(c.Rank.ShortString() + c.Suit.ShortString())
}

func (c Card) ShortString() string {
	return c.Rank.ShortString() + c.Suit.ShortString()
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}
