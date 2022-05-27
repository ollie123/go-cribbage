package main

type Combination struct {
	n, r       int
	comb       []int
	done       bool
	nextCalled bool
}

func NewCombination(n, r int) *Combination {
	return &Combination{
		n:    n,
		r:    r,
		comb: make([]int, r),
	}
}

func (c *Combination) Comb() []int {
	return c.comb
}

func (c *Combination) Next() bool {
	if c.done {
		return false
	}
	var i int

	// Generate the initial combination on the first call to Next.
	if !c.nextCalled {
		c.nextCalled = true
		for i = 0; i < c.r; i++ {
			c.comb[i] = i
		}
		return true
	}

	// Advance to the next combination on subsequent calls.
	c.done = true
	for i = c.r - 1; i >= 0; i-- {
		if c.comb[i] != i+c.n-c.r {
			c.done = false
			break
		}
	}
	if c.done {
		return false
	}
	c.comb[i]++
	for j := i + 1; j < c.r; j++ {
		c.comb[j] = c.comb[j-1] + 1
	}
	return true
}

func Select[T any](indices []int, x, out []T) []T {
	if len(indices) > len(out) {
		panic("len(indices) > len(out)")
	}
	for i, j := range indices {
		out[i] = x[j]
	}
	return out
}
