package main

import "sort"

type Scorable interface {
	Hand() Hand
	Priority() int
	Score() int
	Type() string
}

// Fifteen represents two or more cards whose ranks add to 15 (Aces low, face cards counting as 10).
type Fifteen struct {
	hand Hand
}

func (f Fifteen) Hand() Hand {
	return f.hand
}

func (f Fifteen) Priority() int {
	return 1
}

func (f Fifteen) Score() int {
	return 2
}

func (f Fifteen) Type() string {
	return "Fifteen"
}

// Flush represents four or five cards with the same suit.
type Flush struct {
	hand Hand
}

func (f Flush) Hand() Hand {
	return f.hand
}

func (f Flush) Priority() int {
	return 0
}

func (f Flush) Score() int {
	return len(f.hand)
}

func (f Flush) Type() string {
	return "Flush"
}

// Nobs represents a Jack in the hand with the same suit as the cut card.
type Nobs struct {
	card Card
}

func (k Nobs) Hand() Hand {
	return Hand{k.card}
}

func (k Nobs) Priority() int {
	return 4
}

func (k Nobs) Score() int {
	return 1
}

func (k Nobs) Type() string {
	return "Nobs"
}

// Pair represents two cards with the same rank.
type Pair struct {
	hand Hand
}

func (p Pair) Hand() Hand {
	return p.hand
}

func (p Pair) Priority() int {
	return 3
}

func (p Pair) Score() int {
	return 2
}

func (p Pair) Type() string {
	return "Pair"
}

// Run represents three or more cards whose ranks are consecutive.
type Run struct {
	hand Hand
}

func (r Run) Hand() Hand {
	return r.hand
}

func (r Run) Priority() int {
	return 2
}

func (r Run) Score() int {
	return len(r.hand)
}

func (r Run) Type() string {
	return "Run"
}

type ScoreCard []Scorable

func (sc ScoreCard) SortByPriority() {
	sort.Slice(sc, func(i, j int) bool {
		if sc[i].Priority() != sc[j].Priority() {
			return sc[i].Priority() < sc[j].Priority()
		}
		return len(sc[i].Hand()) < len(sc[j].Hand())
	})
}

func Score(hand Hand, cut Card) ScoreCard {
	var sc ScoreCard
	handWithCut := hand.Add(cut)
	handWithCut.SortByRank()

	// Nobs
	for _, c := range hand {
		if c.Rank == Jack && c.Suit == cut.Suit {
			sc = append(sc, Nobs{c})
		}
	}

	// Flush
	if handWithCut.IsSameSuit() {
		sc = append(sc, Flush{handWithCut.Copy()})
	} else if hand.IsSameSuit() {
		sc = append(sc, Flush{hand.Copy()})
	}

	// Pairs, runs, fifteens
	var runLength int
	for n := 5; n > 1; n-- {
		c := NewCombination(len(handWithCut), n)
		h := make(Hand, n)
		for c.Next() {
			h = Select(c.Comb(), handWithCut, h)

			// Pair
			if n == 2 && h.IsSameRank() {
				sc = append(sc, Pair{h.Copy()})
			}

			// Run
			if n >= 3 && n <= 5 && (runLength == 0 || runLength == n) && h.IsRun() {
				runLength = len(h)
				sc = append(sc, Run{h.Copy()})
			}

			// Fifteen
			if h.SumRanks() == 15 {
				sc = append(sc, Fifteen{h.Copy()})
			}
		}
	}
	return sc
}
