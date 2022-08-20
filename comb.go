package main

type Combination struct {
	n, r int
	comb []int
}

func (c *Combination) Comb() []int {
	return c.comb
}

func (c *Combination) Next() bool {
	if c.comb == nil {
		c.init()
		return true
	}
	return c.next()
}

func (c *Combination) init() {
	c.comb = make([]int, c.r)
	for i := 0; i < c.r; i++ {
		c.comb[i] = i
	}
}

func (c *Combination) next() bool {
	i, ok := c.nextIndex()
	if !ok {
		return false
	}
	c.comb[i]++
	for j := i + 1; j < c.r; j++ {
		c.comb[j] = c.comb[j-1] + 1
	}
	return true
}

func (c *Combination) nextIndex() (int, bool) {
	for i := c.r - 1; i >= 0; i-- {
		if c.comb[i] != i+c.n-c.r {
			return i, true
		}
	}
	return 0, false
}

func NewCombination(n, r int) *Combination {
	return &Combination{n: n, r: r}
}

func Select[T any](selectors []int, in, out []T) []T {
	if out == nil {
		out = make([]T, len(selectors))
	}
	for i, j := range selectors {
		out[i] = in[j]
	}
	return out
}
