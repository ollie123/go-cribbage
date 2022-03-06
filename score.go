package main

type Scorable interface {
	Hand() Hand
	Order() int
	Score() int
	Type() string
}

type Flush struct {
	hand Hand
}

func (f Flush) Hand() Hand {
	return f.hand
}

func (f Flush) Order() int {
	return 0
}

func (f Flush) Score() int {
	return len(f.hand)
}

func (f Flush) Type() string {
	return "Flush"
}

type Run struct {
	hand Hand
}

func (r Run) Hand() Hand {
	return r.hand
}

func (r Run) Order() int {
	return 2
}

func (r Run) Score() int {
	return len(r.hand)
}

func (r Run) Type() string {
	return "Run"
}

type Total struct {
	hand Hand
}

func (t Total) Hand() Hand {
	return t.hand
}

func (t Total) Order() int {
	return 1
}

func (t Total) Score() int {
	return 2
}

func (t Total) Type() string {
	return "Fifteen"
}

type Pair struct {
	hand Hand
}

func (p Pair) Hand() Hand {
	return p.hand
}

func (p Pair) Order() int {
	return 3
}

func (p Pair) Score() int {
	return 2
}

func (p Pair) Type() string {
	return "Pair"
}

type Nobs struct {
	card Card
}

func (n Nobs) Hand() Hand {
	return Hand{n.card}
}

func (n Nobs) Order() int {
	return 4
}

func (n Nobs) Score() int {
	return 1
}

func (n Nobs) Type() string {
	return "Nobs"
}

type ScoreCard []Scorable

func (sc ScoreCard) Len() int {
	return len(sc)
}

func (sc ScoreCard) Swap(i, j int) {
	sc[i], sc[j] = sc[j], sc[i]
}

func (sc ScoreCard) Less(i, j int) bool {
	if sc[i].Order() != sc[j].Order() {
		return sc[i].Order() < sc[j].Order()
	}
	return len(sc[i].Hand()) < len(sc[j].Hand())
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
		handWithCut.Combs(n, func(h Hand) error {
			if n == 2 && h.IsSameRank() {
				sc = append(sc, Pair{h.Copy()})
			}
			if n >= 3 && n <= 5 && (runLength == 0 || runLength == n) && h.IsRun() {
				runLength = len(h)
				sc = append(sc, Run{h.Copy()})
			}
			if h.SumRanks() == 15 {
				sc = append(sc, Total{h.Copy()})
			}
			return nil
		})
	}
	return sc
}
