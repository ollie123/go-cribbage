package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(?i)^([0-9]{1,2}|[atjqk])([cdhs])$`)

func ParseCard(s string) (Card, error) {
	ss := re.FindStringSubmatch(s)
	if ss == nil {
		return Card{}, fmt.Errorf("invalid card: %v", s)
	}
	rank, err := ParseRank(ss[1])
	if err != nil {
		return Card{}, err
	}
	suit, err := ParseSuit(ss[2])
	if err != nil {
		return Card{}, err
	}
	return Card{rank, suit}, nil
}

func ParseHand(s string) (Hand, error) {
	ss := strings.Fields(s)
	hand := make(Hand, len(ss))
	for i, s := range ss {
		c, err := ParseCard(s)
		if err != nil {
			return nil, err
		}
		hand[i] = c
	}
	return hand, nil
}

func ParseRank(s string) (Rank, error) {
	switch s {
	case "a", "A":
		return Ace, nil
	case "t", "T":
		return Ten, nil
	case "j", "J":
		return Jack, nil
	case "q", "Q":
		return Queen, nil
	case "k", "K":
		return King, nil
	}
	rank, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if rank < 1 || rank > 10 {
		return 0, fmt.Errorf("invalid rank: %v", s)
	}
	return Rank(rank - 1), nil
}

func ParseSuit(s string) (Suit, error) {
	switch s {
	case "c", "C":
		return Clubs, nil
	case "d", "D":
		return Diamonds, nil
	case "h", "H":
		return Hearts, nil
	case "s", "S":
		return Spades, nil
	}
	return 0, fmt.Errorf("invalid suit: %v", s)
}
