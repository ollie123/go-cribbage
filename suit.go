package main

import "github.com/fatih/color"

type Suit int

func (s Suit) Color() *color.Color {
	return suits[s].color
}

func (s Suit) ShortString() string {
	return suits[s].shortString
}

func (s Suit) String() string {
	return suits[s].string
}

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
	numSuits

	NumSuits = int(numSuits)
)

type suit struct {
	color               *color.Color
	shortString, string string
}

var suits = [numSuits]suit{
	{color.New(color.FgHiWhite), "♣", "Clubs"},
	{color.New(color.FgRed), "♦", "Diamonds"},
	{color.New(color.FgRed), "♥", "Hearts"},
	{color.New(color.FgHiWhite), "♠", "Spades"},
}
