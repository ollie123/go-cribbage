package main

import (
	"math/rand"
	"sort"
	"strings"
)

type Hand []Card

func (h Hand) Add(c ...Card) Hand {
	hand := make(Hand, 0, len(h)+len(c))
	hand = append(hand, h...)
	hand = append(hand, c...)
	return hand
}

func (h Hand) ColorString() string {
	ss := make([]string, len(h))
	for i, c := range h {
		ss[i] = c.ColorString()
	}
	return strings.Join(ss, ", ")
}

func (h Hand) Copy() Hand {
	hand := make(Hand, 0, len(h))
	hand = append(hand, h...)
	return hand
}

func (h Hand) CountByRank() []int {
	counts := make([]int, NumRanks)
	for _, c := range h {
		counts[c.Rank]++
	}
	return counts
}

func (h Hand) CountBySuit() []int {
	counts := make([]int, NumSuits)
	for _, c := range h {
		counts[c.Suit]++
	}
	return counts
}

func (h Hand) IsRun() bool {
	for i := 1; i < len(h); i++ {
		if h[i].Rank != h[i-1].Rank+1 {
			return false
		}
	}
	return true
}

func (h Hand) IsSameRank() bool {
	for i := 1; i < len(h); i++ {
		if h[i].Rank != h[i-1].Rank {
			return false
		}
	}
	return true
}

func (h Hand) IsSameSuit() bool {
	for i := 1; i < len(h); i++ {
		if h[i].Suit != h[i-1].Suit {
			return false
		}
	}
	return true
}

func (h Hand) Shuffle() {
	rand.Shuffle(len(h), func(i, j int) {
		h[i], h[j] = h[j], h[i]
	})
}

func (h Hand) SortByRank() {
	sort.Slice(h, func(i, j int) bool {
		if h[i].Rank != h[j].Rank {
			return h[i].Rank < h[j].Rank
		}
		return h[i].Suit < h[j].Suit
	})
}

func (h Hand) SortBySuit() {
	sort.Slice(h, func(i, j int) bool {
		if h[i].Suit != h[j].Suit {
			return h[i].Suit < h[j].Suit
		}
		return h[i].Rank < h[j].Rank
	})
}

func (h Hand) SumRanks() int {
	var sum int
	for _, c := range h {
		sum += c.Rank.Score()
	}
	return sum
}

func SortedByRank(h Hand) Hand {
	hand := h.Copy()
	hand.SortByRank()
	return hand
}

func SortedBySuit(h Hand) Hand {
	hand := h.Copy()
	hand.SortBySuit()
	return hand
}
